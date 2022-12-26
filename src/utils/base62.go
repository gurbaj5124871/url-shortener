package utils

import (
	"fmt"
	"math"
	"regexp"
)

const (
	base            = 62
	digitOffset     = 48
	lowercaseOffset = 61
	uppercaseOffset = 55
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func char2ord(char string) (int, error) {
	if matched, _ := regexp.MatchString("[0-9]", char); matched {
		return int([]rune(char)[0] - digitOffset), nil
	} else if matched, _ := regexp.MatchString("[A-Z]", char); matched {
		return int([]rune(char)[0] - uppercaseOffset), nil
	} else if matched, _ := regexp.MatchString("[a-z]", char); matched {
		return int([]rune(char)[0] - lowercaseOffset), nil
	} else {
		return -1, fmt.Errorf("%s is not a valid character", char)
	}
}

func ord2char(ord int64) (string, error) {
	switch {
	case ord < 10:
		return string(rune(ord + digitOffset)), nil
	case ord >= 10 && ord <= 35:
		return string(rune(ord + uppercaseOffset)), nil
	case ord >= 36 && ord < 62:
		return string(rune(ord + lowercaseOffset)), nil
	default:
		return "", fmt.Errorf("%d is not a valid integer in the range of base %d", ord, base)
	}
}

func Base62Decode(str string) (*int64, error) {
	var sum int64 = 0
	for i, c := range reverse(str) {
		if d, err := char2ord(string(c)); err == nil {
			sum = sum + int64(d*int(math.Pow(base, float64(i))))
		} else {
			return nil, err
		}
	}

	return &sum, nil
}

func Base62Encode(digits int64) (string, error) {
	if digits == 0 {
		return "0", nil
	}

	str := ""
	for digits >= 0 {
		remainder := digits % base
		if s, err := ord2char(remainder); err != nil {
			return "", err
		} else {
			str = s + str
		}

		if digits == 0 {
			break
		}
		digits = int64(digits / base)
	}
	return str, nil
}
