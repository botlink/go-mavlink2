package mavlink2

/*
Generated using mavgen - https://github.com/ArduPilot/pymavlink/

Copyright 2020 queue-b <https://github.com/queue-b>

Permission is hereby granted, free of charge, to any person obtaining a copy
of the generated software (the "Generated Software"), to deal
in the Generated Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Generated Software, and to permit persons to whom the Generated
Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Generated Software.

THE GENERATED SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE GENERATED SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE GENERATED SOFTWARE.
*/

import "encoding/binary"

// X25CRC represents a 2-byte CRC redundancy check used to check MAVLink frame integrity
type X25CRC struct {
	value uint16
}

func (hash *X25CRC) Write(input []byte) (int, error) {
	accum := hash.value

	for _, b := range input {
		tmp := b ^ byte(accum&0xff)
		tmp = (tmp ^ (tmp << 4)) & 0xff
		accum = (accum >> 8) ^ (uint16(tmp) << 8) ^ (uint16(tmp) << 3) ^ (uint16(tmp) >> 4)
	}

	hash.value = accum

	return len(input), nil
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (hash *X25CRC) Sum(b []byte) []byte {
	crcBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(crcBytes, hash.value)
	return append(b, crcBytes...)
}

// Reset resets the Hash to its initial state.
func (hash *X25CRC) Reset() {
	hash.value = 0xffff
}

// Size returns the number of bytes Sum will return.
func (hash *X25CRC) Size() int {
	return 2
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (hash *X25CRC) BlockSize() int {
	return 1
}
