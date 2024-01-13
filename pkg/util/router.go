package util

import "regexp"

func MatchAndHandle(input string, pattern string) bool {
	match, _ := regexp.MatchString(pattern, input)
	return match
}
