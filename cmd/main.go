package main

import (
	"fmt"
	"io"
	"net"
	"sync"

	mavlink2 "github.com/queue-b/go-mavlink2"
)

func main() {
	udp, err := net.ListenPacket("udp4", "0.0.0.0:14551")

	if err != nil {
		fmt.Println(err)
		return
	}

	readPipe, writePipe := io.Pipe()

	var wg sync.WaitGroup

	stream := mavlink2.MAVLinkFrameStream{}

	errors := make(chan error)
	frames := make(chan mavlink2.Frame)

	stream.Errors = errors
	stream.Frames = frames

	stream.Run(readPipe, nil)

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
