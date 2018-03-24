package strings

import (
	"fmt"
	"strings"
)

/*

Problem Description :

given a number 'N' print all unique combinations of vowels of length n that are in alphabetical order.

 Eg:
 for N = 1, Output => a, e, i, o, u
 for N = 2, Output => aa, ae, ai, ao, au, ee, ei, eo, eu, ii, io, iu, oo, ou, uu
*/

var vowels = []string{"a", "e", "i", "o", "u"}

func BeautifulStrings(n int, prefix string) {

	if n == 0 {
		if isAlphabetical(prefix) {
			fmt.Print(prefix, ",")
		}
		return
	}

	for i := 0; i < 5; i++ {
		newPrefix := prefix + vowels[i]
		BeautifulStrings(n-1, newPrefix)
	}

}

func isAlphabetical(input string) bool {
	stringArr := strings.Split(input, "")

	for i := 0; i < len(stringArr)-1; i++ {
		if strings.Compare(stringArr[i], stringArr[i+1]) > 0 {
			return false
		}
	}
	return true
}
