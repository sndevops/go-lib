package lib

import (
	"errors"
	"net"
	"os"
	"strconv"
	"unicode/utf8"
	"math/rand"
	"time"
)

var num = []rune("0123456789")
var lenNum = len(num)

var alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lenAlpha = len(alpha)

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = num[rand.Intn(lenNum)]
	}
	return string(b)
}

func randAlpha(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = alpha[rand.Intn(lenAlpha)]
	}
	return string(b)
}

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

//http://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetIPAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
