package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/common"
)

func main() {
	udp, err := net.ListenPacket("udp4", "0.0.0.0:14551")

	if err != nil {
		fmt.Println(err)
		return
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	dialects := mavlink2.Dialects([]mavlink2.Dialect{common.DialectCommon{}})

	wg.Add(1)
	go func() {
	runloop:
		for {
			select {
			case <-sigs:
				break runloop
			default:
				packet := make([]byte, 1024)
				length, _, err := udp.ReadFrom(packet)

				if err != nil {
					fmt.Println(err)
					break runloop
				}

				frame, err := mavlink2.FrameFromBytes(packet[:length])

				if err != nil {
					fmt.Println(err)
					continue
				}

				err = dialects.Validate(frame)

				if err != nil && err != mavlink2.ErrUnknownMessage {
					fmt.Printf("*** Error validating %v\n", err)
					fmt.Println(frame)
				}

				message, err := dialects.GetMessage(frame)

				if err != nil {
					fmt.Printf("*** Error parsing %v\n", err)
				}

				fmt.Println(message)
			}
		}

		udp.Close()
		wg.Done()
	}()

	wg.Wait()
}
