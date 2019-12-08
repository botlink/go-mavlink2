package util

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// ErrNoAddress indicates that the ReadWriter has not read a packet from a remote source
// and cannot determine where to write packets to
var ErrNoAddress = errors.New("No address to write to")

// UDPReadWriteCloser implements io.ReadWriter semantics on top of a UDP packet connection
type UDPReadWriteCloser struct {
	sync.Mutex
	lastAddr *net.UDPAddr
	udp      net.PacketConn
}

// NewUDPReadWriteCloser creates a new instance of UDPReadWriter
func NewUDPReadWriteCloser(address string, port int) (*UDPReadWriteCloser, error) {
	packets, err := net.ListenPacket("udp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return nil, err
	}

	stream := &UDPReadWriteCloser{
		udp: packets,
	}

	return stream, nil
}

func (stream *UDPReadWriteCloser) Read(p []byte) (int, error) {
	count, addr, err := stream.udp.ReadFrom(p)

	if err != nil {
		return count, err
	}

	stream.Lock()
	defer stream.Unlock()
	if stream.lastAddr == nil {
		if localAddr, ok := addr.(*net.UDPAddr); ok {
			stream.lastAddr = localAddr
		}
	}

	return count, err
}

func (stream *UDPReadWriteCloser) Write(p []byte) (int, error) {
	stream.Lock()
	defer stream.Unlock()
	if stream.lastAddr == nil {
		return 0, ErrNoAddress
	}

	return stream.udp.WriteTo(p, stream.lastAddr)
}

// Close closes the underlying UDP connection
func (stream *UDPReadWriteCloser) Close() error {
	return stream.udp.Close()
}
