package main

import (
	"flag"
	"fmt"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "", "file name")
}

func main() {
	opts := flag.Bool("opts", false, "limit of bytes to copy")
	tags := flag.Bool("tags", false, "offset in input file")
	flag.Parse()

	fmt.Printf("%+v\t%+v\t%+v\n", file, *opts, *tags)

	if file == "" {
		fmt.Println("flags -f must be specified")
		return
	}
}
