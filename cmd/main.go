package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/common"
)

func main() {
	udp, err := net.ListenPacket("udp4", "0.0.0.0:14551")

	if err != nil {
		fmt.Println(err)
		return
	}

	readPipe, writePipe := io.Pipe()

	var wg sync.WaitGroup

	dialects := mavlink2.Dialects([]mavlink2.Dialect{common.DialectCommon{}})

	stream := mavlink2.MAVLinkStream{}

	messages := make(chan mavlink2.Message)
	frames := make(chan mavlink2.Frame)

	stream.Dialects = dialects
	stream.Messages = messages
	stream.Frames = frames

	stream.Run(readPipe)

	wg.Add(2)
	go func() {
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			packet := make([]byte, 1024)
			length, _, err := udp.ReadFrom(packet)

			if err != nil {
				fmt.Println(err)
				break
			}

			writePipe.Write(packet[:length])
		}
	}()

	go func() {
		for {
			select {
			case frame := <-frames:
				fmt.Println(frame)
			}
		}
	}()

	wg.Wait()
}
