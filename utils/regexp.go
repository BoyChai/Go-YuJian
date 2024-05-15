package utils

import (
	"regexp"
)

// URL校验
func CheckURL(s string) bool {
	reg := regexp.MustCompile(`^(https?://)?([\da-z\.-]+)\.([a-z\.]{2,6})([/\w \.-]*)*/?$`)
	return reg.MatchString(s)
}

// 纯数字校验
func CheckDigits(s string) bool {
	reg := regexp.MustCompile(`^\d+$`)
	return reg.MatchString(s)
}
