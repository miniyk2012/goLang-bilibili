package homework

import "strings"

func judgePalindrome(s string) bool {
	runeS := []rune(strings.ToUpper(s))
	start, end := 0, len(runeS)-1
	for end > start {
		if runeS[start] != runeS[end] {
			return false
		}
		end--
		start++
	}
	return true
}
