package mavlink2_test

import (
	"bytes"
	"context"
	"testing"
	"time"

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

	t.Run("Mavlink2-BigFrame", func(t *testing.T) {
		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(mav2BigFrame)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStream(buf, inputFrames, dialects, false)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyBigFrame(recvd, t)
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

var mav2BigFrame = []byte{0xfd, 0xfd, 0x29, 0x00, 0x00, 0xf7, 0x01, 0x01, 0xe6, 0x00, 0x00, 0xa7, 0xe3, 0x66, 0x5c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x7f, 0x00, 0x00, 0xc0, 0x7f, 0x1d, 0xca, 0xf6, 0x3b, 0x6c, 0xb0, 0x94, 0x3b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x8b, 0x98, 0x3e, 0xbc, 0x6e, 0x6c, 0x3f, 0xe5, 0xc7, 0xb4, 0xfd, 0x20, 0x00, 0x00, 0xf8, 0x01, 0x01, 0x8d, 0x00, 0x00, 0xf7, 0x61, 0x67, 0x5c, 0x03, 0x00, 0x00, 0x00, 0x88, 0x7c, 0x9f, 0x43, 0x88, 0x7c, 0x9f, 0x43, 0x45, 0xba, 0xb5, 0x40, 0x45, 0xba, 0xb5, 0x40, 0xf4, 0xdc, 0xb1, 0x40, 0x40, 0x54, 0xf7, 0x3d, 0xf6, 0x2f, 0xfd, 0x02, 0x00, 0x00, 0xf9, 0x01, 0x01, 0xf5, 0x00, 0x00, 0x00, 0x01, 0x38, 0xf1, 0xfd, 0x14, 0x00, 0x00, 0xfa, 0x01, 0x01, 0x24, 0x00, 0x00, 0xf0, 0x31, 0x67, 0x5c, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xb6, 0x03, 0xaa, 0x99, 0xfd, 0x24, 0x00, 0x00, 0xfb, 0x01, 0x01, 0x01, 0x00, 0x00, 0x2c, 0xc0, 0x28, 0xd2, 0x2c, 0x40, 0x08, 0xd2, 0x0f, 0x80, 0x0e, 0xc2, 0xb1, 0x02, 0xb8, 0xad, 0x4a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1d, 0x02, 0x00, 0x00, 0x00, 0x02, 0x69, 0x46, 0xfd, 0x2a, 0x00, 0x00, 0xfc, 0x01, 0x01, 0x64, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x33, 0x39, 0x33, 0x31, 0x33, 0x38, 0x35, 0x31, 0x31, 0x35, 0x30, 0x30, 0x32, 0x61, 0x30, 0x30, 0x32, 0x64, 0xbb, 0xa7, 0xfd, 0x35, 0x00, 0x00, 0xfd, 0x01, 0x01, 0x65, 0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7a, 0xc4, 0x00, 0x00, 0x7a, 0xc4, 0x45, 0xba, 0xb5, 0x40, 0xcd, 0xec}

func verifyBigFrame(frame mavlink2.Frame, t *testing.T) {
	if frame.GetVersion() != 2 {
		t.Errorf("Expected version 2, got %d", frame.GetVersion())
	}

	if frame.GetMessageLength() != 253 {
		t.Errorf("Expected length 253, got %d", frame.GetMessageLength())
	}

	if frame.GetMessageSequence() != 0 {
		t.Errorf("Expected sequence 2, got %d", frame.GetMessageSequence())
	}

	if frame.GetMessageID() != 0x00e601 {
		t.Errorf("Expected message ID 0, got %d", frame.GetMessageID())
	}

	if frame.GetChecksum() != ((121 << 8) | 201) {
		t.Errorf("Expected checksum %d, got %d", ((121 << 8) | 201), frame.GetChecksum())
	}

	dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}

	if err := dialects.Validate(frame); err != nil {
		t.Errorf("Invalid packet: %v", err)
	}
}

type testLogger struct {
	bufRead    *bytes.Buffer
	bufWritten *bytes.Buffer
}

func (t testLogger) LogRead(msg []byte) error {
	t.bufRead.Write(msg)
	return nil
}

func (t testLogger) LogWritten(msg []byte) error {
	t.bufWritten.Write(msg)
	return nil
}

func TestFrameLogging(t *testing.T) {
	t.Run("LogRead", func(t *testing.T) {
		logger := testLogger{new(bytes.Buffer), new(bytes.Buffer)}
		inputFrames := make(chan mavlink2.Frame, 1)
		buf := bufReadWriteCloser{bytes.NewBuffer(mav2Heartbeat)}
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStreamWithLogger(buf, inputFrames, dialects, false, logger)

		ctx := context.Background()
		go mavlinkStream.ReadContext(ctx)

		recvd, ok := <-mavlinkStream.Read()

		if !ok {
			t.Errorf("Unable to read from stream")
			return
		}

		verifyHeartbeatFrame(recvd, t)
		if bytes.Compare(recvd.Bytes(), logger.bufRead.Bytes()) != 0 {
			t.Errorf("Message read does not match message logged. %v vs %v\n", recvd.Bytes(), logger.bufRead.Bytes())
		}
	})

	t.Run("LogWritten", func(t *testing.T) {
		logger := testLogger{new(bytes.Buffer), new(bytes.Buffer)}
		inputFrames := make(chan mavlink2.Frame, 1)
		msg := mavlink2.FrameV2(mav2Heartbeat)
		buf := bufReadWriteCloser{new(bytes.Buffer)}
		buf.Grow(len(msg))
		dialects := mavlink2.Dialects{common.Dialect{}, ardupilotmega.Dialect{}}
		mavlinkStream := mavlink2.NewFrameStreamWithLogger(buf, inputFrames, dialects, false, logger)

		ctx := context.Background()
		go mavlinkStream.WriteContext(ctx)

		mavlinkStream.Write() <- msg
		// wait for goroutine
		time.Sleep(time.Millisecond * 100)
		if mavlinkStream.GetCloseErr() != nil {
			t.Errorf("error %s\n", mavlinkStream.GetCloseErr().Error())
		}
		if bytes.Compare(msg.Bytes(), logger.bufWritten.Bytes()) != 0 {
			t.Errorf("Message written does not match message logged. %v vs %v\n", msg.Bytes(), logger.bufWritten.Bytes())
		}
	})
}
