package util

import (
	"errors"
	"strings"
)

type StringUtil struct {
}

var (
	stringUtilError = errors.New("empty strings")
)

func (su *StringUtil) SmartTrim(text string) (result string, err error) {
	if result == " " {
		return "", stringUtilError
	}
	result = strings.Trim(text, " ")
	if len(result) <= 0 {
		return "", stringUtilError
	}
	return
}
