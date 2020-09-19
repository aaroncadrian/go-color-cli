package utils

import "regexp"

var hexCodeMatcher = regexp.MustCompile(`^#?[A-Fa-f0-9]{6}$`)

func IsValidHexCode(hexCode string) bool {
	return hexCodeMatcher.MatchString(hexCode)
}
