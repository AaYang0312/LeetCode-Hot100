package main

import (
	"bytes"
	"unicode"
)

func decodeString(s string) string {
	i := 0
	var decode func() (res []byte)
	decode = func() (res []byte) {
		k := 0
		for i < len(s) {
			c := s[i]
			i++
			if unicode.IsLetter(rune(c)) {
				res = append(res, c)
			} else if unicode.IsDigit(rune(c)) {
				k = k*10 + int(c-'0')
			} else if c == '[' {
				t := decode() // 处理递归如abc3[cd]xyz
				res = append(res, bytes.Repeat(t, k)...)
			} else {
				break
			}
		}
		return res
	}
	return string(decode())
}
