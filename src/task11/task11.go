package task11

import (
	"fmt"
	"regexp"
	"strings"
)

func incrementString(s string) string {
	// 97 is "a", 122 is "z"
	for i := len(s) - 1; i >= 0; i-- {
		next := s[i] + 1
		if next > 122 {
			next = 97
		}
		s = s[:i] + string(next) + s[i+1:]
		if (next != 97) {
			break
		}
	}
	return s
}

func incrementIOL(s string) string {
	// 105 = i, 108 = l, 111 = o
	for i := 0; i < len(s); i++ {
		if s[i] == 105 || s[i] == 108 || s[i] == 111 {
			return s[:i] + string(s[i] + 1) + strings.Repeat("a", len(s[i+1:]))
		}
	}
	return s
}

func isValidPassword(s string) bool {
	iol, _ := regexp.Compile(`[iol]`)
	if (iol.MatchString(s)) {
		return false
	}

	hasThreeSeq := false
	hasTwoPairs := false

	for i := 0; i < len(s); i++ {
		if !hasThreeSeq && i > 1 && s[i] - s[i-1] == 1 && s[i-1] - s[i-2] == 1 {
			hasThreeSeq = true
		}

		for j := 0; j < i - 2; j++ {
			if !hasTwoPairs && s[j+1] == s[j] && s[i] == s[i-1] {
				hasTwoPairs = true
			}
		}
	}

	if (!hasThreeSeq) {
		return false
	}

	if (!hasTwoPairs) {
		return false
	}

	return true
}

func nextPassword(s string) string {
	for {
		s = incrementIOL(incrementString(s))
		if isValidPassword(s) {
			break
		}
	}

	return s
}

func Solve() {
	fmt.Println("Next password is", nextPassword("hepxcrrq"))
	fmt.Println("Next password is", nextPassword("hepxxyzz"))
}