package main

import (
	"context"
	"log"
	"sync"

	mavlink2 "github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/ardupilotmega"
	"github.com/queue-b/go-mavlink2/common"
	"github.com/queue-b/go-mavlink2/util"
)

func main() {
	rwc, err := util.NewUDPReadWriteCloser("0.0.0.0", 14551)

	if err != nil {
		log.Fatal(err)
	}

	input := make(chan mavlink2.Frame)

	dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}

	stream := mavlink2.NewFrameStream(rwc, input, dialects)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	go stream.ReadContext(ctx)
	go stream.WriteContext(ctx)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			case frame, ok := <-stream.Read():
				if !ok {
					log.Fatal("Bad frame read")
				}

				msg, err := dialects.GetMessage(frame)

				if err != nil {
					log.Fatal(err)
				}

				log.Println(msg)
			}
		}
	}()

	wg.Wait()
	cancel()
}
