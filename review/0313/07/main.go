package main

import "fmt"

func findWordConcatenation(str string, words []string) []int {
	if len(str) == 0 || len(words) == 0 {
		return []int{}
	}

	var charFrequencyMap = make(map[string]int, len(words))
	for _, word := range words {
		charFrequencyMap[word]++
	}

	var startIndices []int
	var wordCount = len(words)
	var wordLength = len(words[0])

	for i := 0; i <= len(str)-wordCount*wordLength; i++ {
		var wordSeen = make(map[string]int)
		for j := 0; j < wordCount; j++ {
			var nextWordIndex = i + j*wordLength
			word := str[nextWordIndex : nextWordIndex+wordLength]

			if _, ok := charFrequencyMap[word]; !ok {
				break
			}

			wordSeen[word]++
			if wordSeen[word] > charFrequencyMap[word] {
				break
			}

			if j+1 == wordCount {
				startIndices = append(startIndices, i)
			}
		}
	}

	return startIndices
}

func main() {
	str := "catfoxcat"
	words := []string{"cat", "fox"}

	str2 := "catcatfoxfox"
	words2 := []string{"cat", "fox"}

	fmt.Println(findWordConcatenation(str, words))
	fmt.Println(findWordConcatenation(str2, words2))
}
