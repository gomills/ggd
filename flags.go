package main

import "flag"

// Flags represents the expected CLI flags
type Flags struct {
	columnsN int  // number of columns with which to print the hex dump
	bfSize   int  // buffer size in MiB for the stdin reader
	rainbow  bool // wether to print the hex dump with rainbow colors or not
}

// parseFlags parses CLI flags and returns a Flags instance exposing them
func parseFlags() Flags {

	c := flag.Int("c", defaultNColumns, "Number of columns in which to display the hex dump. Default: 8 columns.")
	bf := flag.Int("bf", defaultReadingBuffer, "Buffer (MiB) size to assignate to the buffered reader. Default: 1 MiB")
	rainbow := flag.Bool("color", false, "Choose wether to print a fully rainbowed hex dump. Default: false :(")

	flag.Parse()

	return Flags{columnsN: *c, bfSize: *bf, rainbow: *rainbow}
}
