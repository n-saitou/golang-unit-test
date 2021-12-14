package my_strings

import (
	"strings"
)

func Compare(a, b string) int {
	if a < b {
		return -1
	}

	if a > b {
		return 1
	}

	return 0
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
	var found bool
	sliceS := strings.Split(s, "")
	sliceSubstr := strings.Split(substr, "")
	lenS := len(sliceS)
	lenSubstr := len(sliceSubstr)

	if lenS < lenSubstr {
		return false
	}

	matchCounter := 0
	for i := 0; i < lenS; i++ {
		for i2 := matchCounter; i2 < lenSubstr; i2++ {
			isMatch := sliceS[i] == sliceSubstr[i2]
			isLast := i2 == lenSubstr-1

			if isMatch {
				matchCounter++

				if isLast {
					found = true
				}
				break
			}

			if isLast {
				matchCounter = 0
			}
		}

		if found {
			break
		}
	}

	return found
}

func Count(s, substr string) int {
	var counter int
	for _, c := range s {
		if c == rune(substr[0]) {
			counter++
		}
	}

	return counter
}

func Fields(s string) []string {
	res := []string{""}
	strings.Trim(s, " ")
	sliceS := strings.Split(strings.Trim(s, " "), "")
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] == " " {
			if res[len(res)-1] == "" {
				res = res[:len(res)-1]
			}
			res = append(res, "")
		} else {
			res[len(res)-1] += sliceS[i]
		}
	}
	return res
}
