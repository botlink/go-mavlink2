package mavlink2_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/botlink/go-mavlink2"
	"github.com/botlink/go-mavlink2/ardupilotmega"
	"github.com/botlink/go-mavlink2/common"
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
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyHeartbeatFrame(recvd, t)
	})

	t.Run("Mav2HeartbeatBadCrc", func(t *testing.T) {
		var msg []byte
		msg = append(msg, mav2Heartbeat...)
		msg[len(msg)-1] = msg[len(msg)-1] + 1
		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()

		// Should have received nothing as no valid message in
		// mavlinkStream's buffer.
		if ok {
			t.Errorf("Received message when expecting no message due to bad CRC")
		}
	})

	t.Run("FalseStartByte", func(t *testing.T) {
		// Prepend an extra start byte (and length and compatibility
		// bytes that won't cause us to read past the end of buf)
		msg := []byte{253, 9, 0}
		msg = append(msg, mav2Heartbeat...)

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyHeartbeatFrame(recvd, t)
	})

	t.Run("MultipleHeartbeats", func(t *testing.T) {
		msg := []byte{}
		msg = append(msg, mav2Heartbeat...)
		junk := []byte{253, 9, 0}
		msg = append(msg, junk...)
		msg = append(msg, mav2Heartbeat...)
		msg = append(msg, mav2Heartbeat...)

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		for i := 0; i < 3; i++ {
			recvd, ok := <-mavlinkStream.Read()

			if !ok {
				t.Errorf("Unable to read from stream")
				return
			}

			verifyHeartbeatFrame(recvd, t)
		}
	})

	t.Run("TruncatedZeroBytes", func(t *testing.T) {
		// Generated using pymavlink, system time message
		msg := []byte{253, 1, 0, 0, 3, 255, 0, 2, 0, 0, 3, 230, 132}

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()
		if !ok {
			t.Errorf("Unable to read from stream")
		}
	})

	t.Run("Mavlink1", func(t *testing.T) {
		// Generated using pymavlink, system time message
		msg := []byte{254, 12, 0, 255, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 177, 129}

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()
		if !ok {
			t.Errorf("Unable to read from stream")
		}
	})

	t.Run("Mavlink1BadCrc", func(t *testing.T) {
		// Generated using pymavlink, system time message, 1 added to last byte to invalidate CRC
		msg := []byte{254, 12, 0, 255, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 177, 130}

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()
		// Should have received nothing as no valid message in
		// mavlinkStream's buffer.
		if ok {
			t.Errorf("Received message when expecting no message due to bad CRC")
		}
	})

	t.Run("Mavlink1-ReturnInvalidFrames", func(t *testing.T) {
		// Generated using pymavlink, system time message, 1 added to last byte to invalidate CRC
		msg := []byte{254, 12, 0, 255, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 177, 130}

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, true)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()
		if !ok {
			t.Errorf("Unable to read from stream")
		}
	})

	t.Run("Mav2HeartbeatBadCrc-ReturnInvalidFrames", func(t *testing.T) {
		var msg []byte
		msg = append(msg, mav2Heartbeat...)
		// modify CRC to make invalid
		msg[len(msg)-1] = msg[len(msg)-1] + 1
		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(msg)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, true)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		_, ok := <-mavlinkStream.Read()

		// Should have still received a frame as we created FrameStream
		// that returns invalid frames.
		if !ok {
			t.Errorf("Did not receive invalid frame when FrameStream configured to return invalid frames.")
		}
	})

	t.Run("Mavlink1and2-ReturnInvalidFrames", func(t *testing.T) {
		var msgs [][]byte
		msgs = append(msgs, mav2Heartbeat)
		junk := []byte{253, 9, 0}
		msgs = append(msgs, junk)
		msgs = append(msgs, mav2Heartbeat)
		// Mavlink 1 system time message
		msgs = append(msgs, []byte{254, 12, 0, 255, 0, 2, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 177, 129})
		// Change message ID of Mavlink 1 system time message to "corrupt" message (2 -> 123)
		msgs = append(msgs, []byte{254, 12, 0, 255, 0, 123, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 177, 129})
		msgs = append(msgs, mav2Heartbeat)

		var streamBuffer []byte
		for i := 0; i < len(msgs); i++ {
			streamBuffer = append(streamBuffer, msgs[i]...)
		}

		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(streamBuffer)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, true)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		for i := 0; i < len(msgs); i++ {
			recvd, ok := <-mavlinkStream.Read()

			if !ok {
				t.Errorf("Unable to read from stream")
				return
			}

			expectedLength := msgs[i][1]
			actualLength := recvd.GetMessageLength()
			if actualLength != expectedLength {
				t.Errorf("Did not get expected message length. Actual %d vs expected %d",
					actualLength, expectedLength)
			}
		}
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
