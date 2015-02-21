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

	i, err := strconv.Atoi(s)

	if err != nil {
		return i
	} else {
		return 0 //error
	}
}

//https://gist.github.com/yuchan/2857491
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//http://stackoverflow.com/questions/17602811/check-for-vowels-or-numbers-in-a-string-using-golang
//check if all characters in s are all numeric
func IsNumeric(s string) bool {

	i := 0
	for _, value := range s {
		switch {
		case value >= '0' && value <= '9':
			i++
		}
	}

	//fmt.Println("number of chars: " + strconv.Itoa(utf8.RuneCountInString(s)))

	return utf8.RuneCountInString(s) == i
}

//http://stackoverflow.com/questions/12518876/how-to-check-if-a-file-exists-in-go
// Exists reports whether the named file or directory exists.
func IsFileExist(name string) bool {
	_, err := os.Stat(name)

	return !os.IsNotExist(err)
}
