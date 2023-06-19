package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runes := []rune(s)

	if len(runes) == 0 {
		return "", nil
	}

	if err := validate(runes); err != nil {
		return "", fmt.Errorf("validate error %w", err)
	}

	return unpack(runes), nil
}

func validate(runes []rune) error {
	if unicode.IsDigit(runes[0]) {
		return ErrInvalidString
	}

	str := []byte(string(runes))

	matched, err := regexp.Match(`\d{2,}`, str)
	if err != nil {
		return err
	}

	if matched {
		return ErrInvalidString
	}

	return nil
}

func unpack(runes []rune) string {
	b := strings.Builder{}

	for i, r := range runes {
		current := string(r)

		count := 1
		if i+1 < len(runes) {
			tmp, err := strconv.Atoi(string(runes[i+1]))
			if err == nil {
				count = tmp
			}
		}

		if unicode.IsDigit(r) {
			continue
		}

		b.WriteString(strings.Repeat(current, count))
	}

	return b.String()
}
