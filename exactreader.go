package main

import (
	"bufio"
	"errors"
	"io"
)

// ExactReader implements a composed wrapper over a *bufio.Reader. This is a conveniance
// to call io.ReadFull() over a buffered reader that satisfies io.Reader.
type ExactReader struct {
	*bufio.Reader
	bfSize int
}

// NewExactReader returns a pointer to a newly created ExactReader over an io.Reader with
// a buffer size of bfSize
func NewExactReader(re io.Reader, bfSize int) *ExactReader {
	return &ExactReader{
		Reader: bufio.NewReaderSize(re, bfSize),
		bfSize: bfSize,
	}
}

// ReadFull is a method of ExactReader that calls io.ReadFull() on its *bufio.Reader to read
// all the length of dst.
func (r *ExactReader) ReadFull(dst []byte) (int, error) {

	// return an error if the length to read is bigger than the buffer to warn of a wrongly
	// configured buffer size
	if len(dst) > r.bfSize {
		return 0, errors.New("length of destination is bigger than the reader's buffer")
	}

	// read over ExactReader's *bufio.Reader with io.ReadFull()
	n, err := io.ReadFull(r.Reader, dst)

	return n, err
}
