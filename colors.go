package main

// rainbow ANSI colors bytes for shell printing
var (
	rainbow = [][]byte{
		[]byte("\033[91m"), // bright red
		[]byte("\033[93m"), // bright yellow
		[]byte("\033[92m"), // bright green
		[]byte("\033[96m"), // bright cyan
		[]byte("\033[94m"), // bright blue
		[]byte("\033[95m"), // bright magenta
	}
	ansiReset = []byte("\033[0m")
)

const ansiOpenLen = 5  // len("\033[91m")
const ansiResetLen = 4 // len("\033[0m")

// ByteInd holds a wrapper around a simple integer to keep track of the open file offset, for rainbow color
// printing purposes.
type ByteInd struct {
	int
}

// Increment adds 1 to the respective ByteInd's counter
func (bi *ByteInd) Increment(i int) {
	bi.int++
}

// GetByteInd returns the ByteInd's index
func (bi *ByteInd) GetByteInd() int {
	return bi.int
}
