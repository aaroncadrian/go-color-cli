package utils

import (
	"regexp"
)

var hexCodeMatcher = regexp.MustCompile(`^(?P<Hash>#?)(?P<HexCode>[A-Fa-f0-9]{6})$`)

func IsValidHexCode(hexCode string) bool {
	return hexCodeMatcher.MatchString(hexCode)
}

func ParseHexCode(hexCode string) (string, bool) {
	if !IsValidHexCode(hexCode) {
		return "", false
	}

	result := hexCodeMatcher.FindStringSubmatch(hexCode)

	return result[2], true
}
