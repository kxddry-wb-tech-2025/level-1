package main

import (
	"fmt"
	"unicode"
)

func UniqueCharacters(s string) bool {
	mp := make(map[rune]struct{})
	for _, ch := range s {
		if _, ok := mp[ch]; ok {
			return false
		}
		mp[ch] = struct{}{}
		if unicode.IsLetter(ch) {
			other := ch ^ 0x20
			if _, ok := mp[other]; ok {
				return false
			}
			mp[other] = struct{}{}
		}
	}
	return true
}

func main() {
	fmt.Println(UniqueCharacters("abcd"))
	fmt.Println(UniqueCharacters("abCdefAaf"))
	fmt.Println(UniqueCharacters("aabcd"))
}
