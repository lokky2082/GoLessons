package top

import (
	"fmt"
)

//
func Top(n int, s string) (string, error) {
	fmt.Println("top %d", n)
	return s, nil
}
