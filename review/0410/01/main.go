package main

import "fmt"

func findWordConcatenation(str string, words []string) []int {
	if len(str) == 0 || len(words) == 0 {
		return []int{}
	}

	var wordsMap = make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}

	var resultIndices []int
	wordsCount := len(words)
	singleWordLength := len(words[0])

	for i := 0; i <= len(str)-wordsCount*singleWordLength; i++ {
		var wordsSeenMap = make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*singleWordLength
			word := str[nextWordIndex : nextWordIndex+singleWordLength]

			if _, ok := wordsMap[word]; !ok {
				break
			}

			wordsSeenMap[word]++
			if wordsSeenMap[word] > wordsMap[word] {
				break
			}

			if j+1 == wordsCount {
				resultIndices = append(resultIndices, i)
			}
		}
	}

	return resultIndices
}

func main() {
	str := "catfoxcat"
	words := []string{"cat", "fox"}

	fmt.Println(findWordConcatenation(str, words))
}
