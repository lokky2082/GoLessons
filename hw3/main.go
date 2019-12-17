package main

import (
	"fmt"

	"github.com/lokky2082/GoLessons/tree/master/hw3/unpack"
)

func main() {
	var unpacked = unpack.Unpack("ж5b4")
	fmt.Printf("unpacked = %s\n", unpacked)
}
