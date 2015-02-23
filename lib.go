package lib

import (
	"os"
	"strconv"
	"unicode/utf8"
)

//number of characters in a string, UTF-8 compatible
func StrLen(s string) int {

	return utf8.RuneCountInString(s)
}

func IntToStr(i int) string {

	return strconv.Itoa(i)
}

func StrToInt(s string) int {

	i, _ := strconv.Atoi(s)

	return i
}

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func IsNumeric(s string) bool {

	i := 0
	for _, value := range s {
		switch {
		case value >= '0' && value <= '9':
			i++
		}
	}

	return utf8.RuneCountInString(s) == i
}

func IsFileExist(name string) bool {
	_, err := os.Stat(name)

	return !os.IsNotExist(err)
}
