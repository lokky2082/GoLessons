package main

import (
	"fmt"

	"github.com/lokky2082/GoLessons/tree/master/hw3/unpack"
)

func main() {
	var unpacked, err = unpack.Unpack("/ж5b4")
	if err != nil {
		fmt.Printf("unpackedErr: %s\n", err)
	} else {
		fmt.Printf("unpacked = %s\n", unpacked)
	}
}
