package main

import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	words := strings.Split(s, " ")
	var ret []string
	for _, v := range words {
		word := []byte(v)
		for i:=0; i<len(word)/2;i++{
			word[i], word[len(word)-i-1] = word[len(word)-i-1], word[i]
		}
		ret = append(ret, string(word))
	}

	return strings.Join(ret, " ")
}
