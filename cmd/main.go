package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	mavlink2 "github.com/queue-b/go-mavlink2"
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

	wg.Add(1)
	go func() {
	runloop:
		for {
			select {
			case <-sigs:
				break runloop
			default:
				packet := make([]byte, 1024)
				length, addr, err := udp.ReadFrom(packet)

				fmt.Println(length)
				fmt.Println(addr)

				if err != nil {
					fmt.Println(err)
					break runloop
				}

				frame, err := mavlink2.FrameFromBytes(packet[:length])

				if err != nil {
					fmt.Println(err)
					continue
				}

				fmt.Println(frame)
			}
		}

		udp.Close()
		wg.Done()
	}()

	wg.Wait()
}
