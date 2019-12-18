package unpack

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Unpack convert strings like 'a5b4' into 'aaaaabbbb'
func Unpack(s string) (string, error) {
	var result strings.Builder
	var currentStr = ""
	var count = ""
	fmt.Printf("%s test\n", currentStr)
	for i, r := range s {
		if unicode.IsLetter(r) {
			var countnum, errcount = strconv.Atoi(string(count))
			fmt.Printf("%s is letter\n", string(r))
			if errcount != nil {
				result.WriteString(string(r))
			} else {
				var repeated = strings.Repeat(currentStr, countnum-1)
				result.WriteString(repeated)
				result.WriteString(string(r))
				count = ""
			}
			currentStr = string(r)
		} else if unicode.IsNumber(r) {
			count += string(r)
			fmt.Printf("%s is number\n", count)
			if i == len(s)-1 {
				var countnum, errcount = strconv.Atoi(string(count))
				if errcount != nil {
					fmt.Printf("%s its last number error\n", errcount)
				} else {
					fmt.Printf("%s its last number\n", count)
					var repeated = strings.Repeat(currentStr, countnum-1)
					result.WriteString(repeated)
				}
			}
		} else {
			// как в ошибку передать переменную индекса?
			return "", errors.New("symbol is not alowed at position " + string(i))
		}
	}
	return result.String(), nil
}
