package main

import "fmt"

func findWordConcatenation(str string, words []string) []int {
	if len(str) == 0 || len(words) == 0 {
		return []int{}
	}

	var wordsMap = make(map[string]int, len(words))
	for _, word := range words {
		wordsMap[word]++
	}

	var wordCount = len(words)
	var singleWordLength = len(words[0]) // 单个字符的长度
	var wordIndexes []int

	for i := 0; i <= len(str)-wordCount*singleWordLength; i++ {
		var wordsSeenMap = make(map[string]int)
		for j := 0; j < wordCount; j++ {
			wordIndex := i + j*singleWordLength
			word := str[wordIndex : wordIndex+singleWordLength]

			if _, ok := wordsMap[word]; !ok {
				break
			}

			wordsSeenMap[word]++
			if wordsSeenMap[word] > wordsMap[word] {
				break // 已经遍历过的word数量超过了预期
			}

			if j+1 == wordCount {
				wordIndexes = append(wordIndexes, i)
			}
		}
	}

	return wordIndexes
}

func main() {
	str := "catfoxcat"
	words := []string{"cat", "fox"}

	fmt.Println(findWordConcatenation(str, words))
}
