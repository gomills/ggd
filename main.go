package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

const (
	defaultReadingBuffer = 1 // default buffer for the stdin reader in MiB (e.g: 1 is equivalent 1 MiB)
	defaultNColumns      = 8 // default number of columns with which to print the hex dump
)

// buffered writer for stdout
var buffStdout = bufio.NewWriter(os.Stdout)

func main() {

	// 0. deffer buffer write
	defer buffStdout.Flush()

	// 1. parse flags
	flags := parseFlags()

	// 2.1 create a buffered reader for stdin's file descriptor and an index for its byte offset
	stdinReader := NewExactReader(os.Stdin, flags.bfSize<<20)
	bi := ByteInd{}

	// 3. loop for reading buffered stdin, encoding it to hex and writing the dump to stdout.
	// Print in blocks of one line, each composed of three elements: (|decoded byte index| |hex bytes dump| |ASCII string|)
	numDecodedBytesPerLine := flags.columnsN * 2
	numHexBytesPerLine := numDecodedBytesPerLine * 2
	numSpacesPerHexLine := flags.columnsN - 1

	decodedBytes := make([]byte, numDecodedBytesPerLine)            // decoded raw bytes belonging to the same line
	hexBytes := make([]byte, numHexBytesPerLine)                    // hex bytes belonging to the same line
	hexDump := make([]byte, numHexBytesPerLine+numSpacesPerHexLine) // actual hex dump line with whitespaces for columns to print

	for i := 0; ; i = i + numDecodedBytesPerLine {

		// read numDecodedBytesPerLine bytes from stdin
		readBytes, err := stdinReader.ReadFull(decodedBytes)
		if err == io.EOF { // equivalent to readBytes=0
			return
		}

		// encode them
		hex.Encode(hexBytes, decodedBytes)

		// case where read bytes were less than expected. It's either last line or not enough bytes from content to fill specified columns.
		if readBytes < numDecodedBytesPerLine {

			// case: this dump line is the only one of the file, adjust the columns and sizing
			if i == 0 {
				adjustedColumns := (readBytes + 1) / 2
				hexBytes = hexBytes[:readBytes*2]
				hexDump = hexDump[:(adjustedColumns-1)+len(hexBytes)]
			} else {
				// case: this dump line is the last; set missing hex bytes to a whitespace
				for k := readBytes * 2; k < len(hexBytes); k++ {
					hexBytes[k] = byte(' ')
				}
			}

			hexBytesToDump(hexDump, hexBytes)
			writeToStdout(i, &bi, flags.rainbow, readBytes, decodedBytes, hexDump)
			return
		}

		hexBytesToDump(hexDump, hexBytes)
		writeToStdout(i, &bi, flags.rainbow, readBytes, decodedBytes, hexDump)
	}
}

// hexBytesToDump inserts a whitespace every 4 bytes in hexDump and allocates the new whitespaced slice in dst.
func hexBytesToDump(dst []byte, hexDump []byte) {

	for l, m := 0, 0; l < len(dst); l++ {
		if ((l+1)%5 == 0) && (l > 0) {
			dst[l] = byte(' ')
		} else {
			dst[l] = hexDump[m]
			m++
		}
	}

}

// writeToStdout prints the passed hex dump line to stdout
func writeToStdout(i int, bi *ByteInd, rainbow bool, readBytes int, decodedBytes []byte, hexDump []byte) {

	ascii := createAscii(decodedBytes[:readBytes])

	if rainbow {
		rainbowHexDump := hexDumpToRainbow(bi, hexDump, readBytes)
		fmt.Fprintf(buffStdout, "%08x (%04d)  |  %s  |  %s\n", i, i, string(rainbowHexDump), ascii)
	} else {
		fmt.Fprintf(buffStdout, "%08x (%04d)  |  %s  |  %s\n", i, i, string(hexDump), ascii)
	}

}

// hexDumpToRainbow converts slice hexDump into a rainbow and returns the new rainbowed slice. NSFW.
func hexDumpToRainbow(bi *ByteInd, hexDump []byte, numDecodedBytes int) []byte {

	// create a slice buffer where to hold the rainbow hex dump and keep track of how many
	// hex characters we've iterated over
	rainbowBuf := make([]byte, len(hexDump)+numDecodedBytes*(ansiOpenLen+ansiResetLen))
	hexCharCount := 0

	// iterate over both rainbow buffer and hex dump
	for i, j := 0, 0; j < len(hexDump); j++ {

		// if it's space, just copy it to buffer
		if hexDump[j] == ' ' {
			rainbowBuf[i] = hexDump[j]
			i++
			continue
		}

		// if it's a character that needs color prefixed
		if hexCharCount%2 == 0 {
			for _, r := range rainbow[bi.GetByteInd()%len(rainbow)] {
				rainbowBuf[i] = r
				i++
			}
		}

		// copy character
		rainbowBuf[i] = hexDump[j]
		i++
		hexCharCount++

		// if it's a character that needs ansiReset
		if hexCharCount%2 == 0 {
			for _, b := range ansiReset {
				rainbowBuf[i] = b
				i++
			}
			bi.Increment(1)
		}
	}

	return rainbowBuf
}

// createAscii constructs the printable ascii encoding of the raw bytes
func createAscii(src []byte) string {
	ascii := make([]byte, len(src))
	for i := range src {
		if 32 <= uint8(src[i]) && uint8(src[i]) <= 126 {
			ascii[i] = src[i]
		} else {
			ascii[i] = byte('.')
		}
	}
	return string(ascii)
}
