package unpack

import (
	"fmt"
	"strconv"
	"strings"
)

// Unpack convert strings like 'a5b4' into 'aaaaabbbb'
func Unpack(s string) string {
	var result strings.Builder
	var currentrune = ""
	var count = false
	result.WriteString(s)
	for i, r := range s {
		var num, numerr = strconv.Atoi(string(r))
		if numerr != nil && !count {
			currentrune = string(r)
			fmt.Printf("%s current\n", currentrune)
		} else if numerr != nil && count {
			var countnum, err = strconv.Atoi(string(count))
			if err != nil {
				fmt.Printf("%s err on count %d \n", err, count)
			} else {
				var str = strings.Repeat(currentrune, countnum-1)
				fmt.Printf("%s str \n", str)
				result.WriteString(str)
				count = ""
				currentrune = string(r)
			}

		} else if numerr == nil && !count {
			count = string(num)
			fmt.Printf("%d num && count == \n", count)
		} else if numerr == nil && count {
			fmt.Printf("%s count \n", count)
			count = count + string(num)
		}
		fmt.Printf("%s starts at byte position %d\n", string(r), i)
	}
	//return "xx", "error"
	return result.String()
}
