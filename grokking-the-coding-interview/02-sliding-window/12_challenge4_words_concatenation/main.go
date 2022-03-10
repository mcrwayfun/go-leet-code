package main

import "fmt"

/**
题目：Given a string and a list of words, find all the starting indices of
substrings in the given string that are a concatenation of all
the given words exactly once without any overlapping of words.
It is given that all words are of the same length.

描述：给出一个字符串和一个列表的单词，找到所有子串的起始坐标，需要满足：这些子串是
给定的列表单词的准确连接，没有任何重叠，所有给定的词都是相同长度的。

举例：
Example 1:

Input: String="catfoxcat", Words=["cat", "fox"]
Output: [0, 3]
Explanation: The two substring containing both the words are "catfox" & "foxcat".
Example 2:

Input: String="catcatfoxfox", Words=["cat", "fox"]
Output: [3]
Explanation: The only substring containing both the words is "catfox".

模板：
func findWordConcatenation(str string, words []string) []int
*/

/**
解题思路：暴力破解法
1：使用一个wordFrequencyMap来存储单词出现的次数，wordsCount是多少个单词，wordLength是某个单词的长度（
所有单词的长度都是一致的）
2：外层循环遍历 str，内层循环遍历单词列表
	2-1: i <= len(str)-wordsCount*wordLength (不需要遍历完str）
	2-2: 初始化 wordSeen map 来记录当前已经匹配的 word
	2-3: j < wordsCount
		2-2-1: nextWordIndex = i + j * wordLength, 使用的是下标的方式获取 str中对应的子串.
		2-2-2: word = str[nextWordIndex:nextWordIndex+wordLength]
		2-2-3: if !wordFrequencyMap.contains(word) then break
		2-2-4: if wordSeen[word] > wordFrequencyMap[word]，表示word出现的次数比原来的单词列表中出现
	的次数多，可以 break
		2-2-4: j+1 == wordsCount, 则完全匹配，可以加入当前的下标 i

time complexity: O(n*m)
space complexity: O(n+m)
*/
func findWordConcatenation(str string, words []string) []int {
	if len(str) == 0 || len(words) == 0 {
		return []int{}
	}

	var wordFrequencyMap = make(map[string]int, len(words))
	for _, word := range words {
		wordFrequencyMap[word]++
	}

	var resultIndices []int
	wordsCount := len(words)
	wordLength := len(words[0])

	for i := 0; i <= len(str)-wordsCount*wordLength; i++ {
		var wordsSeen = make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength
			word := str[nextWordIndex : nextWordIndex+wordLength]

			if _, ok := wordFrequencyMap[word]; !ok {
				break
			}

			wordsSeen[word]++
			if wordsSeen[word] > wordFrequencyMap[word] {
				break // the word has higher frequency than required
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
