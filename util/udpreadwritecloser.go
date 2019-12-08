/*
Copyright 2019 queue-b <https://github.com/queue-b>

Permission is hereby granted, free of charge, to any person obtaining a copy
of the Software (the "Software"), to deal
in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Generated
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

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
