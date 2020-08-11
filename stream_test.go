package mavlink2_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/queue-b/go-mavlink2"
	"github.com/queue-b/go-mavlink2/ardupilotmega"
	"github.com/queue-b/go-mavlink2/common"
)

type bufReadWriteCloser struct {
	*bytes.Buffer
}

func (b bufReadWriteCloser) Close() error {
	return nil
}

// Generated using pymavlink
var mav2Heartbeat = []byte{253, 9, 0, 0, 2, 255, 0, 0, 0, 0, 0, 0, 0, 0, 18, 8, 0, 0, 3, 201, 121}

func TestReadFrames(t *testing.T) {
	t.Run("Mav2Heartbeat", func(t *testing.T) {
		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(mav2Heartbeat)}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyHeartbeatFrame(recvd, t)
	})

	t.Run("FalseStartByte", func(t *testing.T) {
		// Prepend an extra start byte (and length and compatibility
		// bytes that won't cause us to read past the end of the buffer)
		msg := []byte{253, 9, 0}
		msg = append(msg, mav2Heartbeat...)

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyHeartbeatFrame(recvd, t)
	})
}

func verifyHeartbeatFrame(frame mavlink2.Frame, t *testing.T) {
	if frame.GetVersion() != 2 {
		t.Errorf("Expected version 2, got %d", frame.GetVersion())
	}

	if frame.GetMessageLength() != 9 {
		t.Errorf("Expected length 9, got %d", frame.GetMessageLength())
	}

	if frame.GetMessageSequence() != 2 {
		t.Errorf("Expected sequence 2, got %d", frame.GetMessageSequence())
	}

	if frame.GetMessageID() != 0 {
		t.Errorf("Expected message ID 0, got %d", frame.GetMessageID())
	}

	if frame.GetChecksum() != ((121 << 8) | 201) {
		t.Errorf("Expected checksum %d, got %d", ((121 << 8) | 201), frame.GetChecksum())
	}

	dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}

	if err := dialects.Validate(frame); err != nil {
		t.Errorf("Error verifying packet checksum: %v", err)
	}
}
