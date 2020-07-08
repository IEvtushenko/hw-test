package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prevSymbol string
	if len(s) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", ErrInvalidString
	}

	for _, r := range s {
		str := string(r)
		if number, err := strconv.Atoi(str); err == nil {
			if _, err := strconv.Atoi(prevSymbol); err == nil {
				return "", ErrInvalidString
			}
			if number == 0 {
				str = result.String()[:len(result.String())-len(prevSymbol)]
				result.Reset()
				result.WriteString(str)
			}
			if number > 0 {
				result.WriteString(strings.Repeat(prevSymbol, number-1))
			}
		} else {
			result.WriteString(str)
		}
		prevSymbol = str
	}

	return result.String(), nil
}
