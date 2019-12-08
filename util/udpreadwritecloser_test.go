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
