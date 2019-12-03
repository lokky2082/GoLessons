package main

import (
	"fmt";
	"github.com/beevik/ntp";
	"os"
)

func main () {
	time,err:= ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err == nil {
		fmt.Println(time)
		os.Exit(0)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}