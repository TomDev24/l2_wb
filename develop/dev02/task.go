package main

import (
	"fmt"
	"unicode"
	"strings"
)

func writeToStr(sb *strings.Builder, c rune, count int) error {
	sb.Grow(count)
	for i := 0; i < count; i++ {
		_, err := sb.WriteRune(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// empty string, means there was an error
func unpack(str string) string {
	var sb strings.Builder
	prev := '-'

	for _, c := range str {
		switch {
		case unicode.IsDigit(c):
			if prev == '-' {
				return ""
			}
			err := writeToStr(&sb, prev, int(c)-'0'-1)
			if err != nil{
				return ""
			}
			prev = '-'

		case unicode.IsLetter(c):
			err := writeToStr(&sb, c, 1)
			if err != nil{
				return ""
			}
			prev = c
		}
	}

	return sb.String()
}

func main(){
	s := unpack("a4bc2d5e3")
	fmt.Println("a4bc2d5e3", s)
}
