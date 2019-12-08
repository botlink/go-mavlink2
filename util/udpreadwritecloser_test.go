package util

import (
	"fmt"
	"net"
	"testing"
)

func TestNewUDPReadWriteCloser(t *testing.T) {
	t.Run("WithBasicParameters", func(t *testing.T) {
		rwc, err := NewUDPReadWriteCloser("0.0.0.0", 5000)

		if err != nil {
			t.Fatalf("Error creating UDPReadWriteCloser - %v\n", err)
		}

		defer rwc.Close()
	})

	t.Run("ReturnsErrorFromNetListenPacket", func(t *testing.T) {
		_, err := net.ListenPacket("udp", fmt.Sprintf("%v:%v", "0.0.0.0", 5000))

		if err != nil {
			t.Fatal("Error listening UDP", err)
		}

		_, err = NewUDPReadWriteCloser("0.0.0.0", 5000)

		if err == nil {
			t.Fatal("NewUDPReadWriteCloser did not return underlying error")
		}
	})
}
