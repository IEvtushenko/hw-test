package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result string
	var prevRune string
	// Попробовать привести к строке
	// Попробовать привести к инту, если так то ...
	// Как взять нужный элемент?
	if len(s) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", ErrInvalidString
	}

	for _, r := range s {
		str := string(r)
		if number, err := strconv.Atoi(str); err == nil {
			if _, err := strconv.Atoi(prevRune); err == nil {
				return "", ErrInvalidString
			}
			if number == 0 {
				result = result[:len(result)-1]
			}
			if number > 0 {
				result += strings.Repeat(string(prevRune), number-1)
			}
		} else {
			result += str
		}
		prevRune = str
	}

	return result, nil
}
