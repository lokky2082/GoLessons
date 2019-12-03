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
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}