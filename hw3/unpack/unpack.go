package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Unpack strings like 'a5b4' into 'aaaaabbbb'
func Unpack(s string) (string, error) {
	var result strings.Builder
	var currentStr = ""
	var count = ""
	for i, r := range s {
		if unicode.IsLetter(r) {
			var countnum, errcount = strconv.Atoi(string(count))
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
			if i == len(s)-1 {
				var countnum, errcount = strconv.Atoi(string(count))
				if errcount != nil {
					return "", errors.New("symbol is not alowed at position " + string(i))
				}
				var repeated = strings.Repeat(currentStr, countnum-1)
				result.WriteString(repeated)
			}
		} else {
			// как в ошибку передать переменную i?
			return "", errors.New("symbol is not alowed at position " + string(i))
		}
	}
	return result.String(), nil
}
