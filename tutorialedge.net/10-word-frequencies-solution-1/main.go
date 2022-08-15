package main

import (
	"fmt"
	"sort"
	"strings"
)

type Word struct {
	Value string
	Score int
}

func CountWords(text string) map[string]int {
	var xs []string
	xs = strings.Split(text, " ")

	wordsMap := make(map[string]int, len(xs))

	for _, w := range xs {
		word := strings.ToLower(w)
		word = strings.Replace(word, ",", "", -1)
		word = strings.Replace(word, ".", "", -1)
		wordsMap[word]++
	}

	return wordsMap
}

func Top5Words(wordmap map[string]int) []Word {
	keys := make([]string, 0, len(wordmap))

	var out []Word

	for k := range wordmap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(x, y int) bool {
		return wordmap[keys[x]] > wordmap[keys[y]]
	})

	for i, k := range keys {
		j := Word{
			Value: k,
			Score: wordmap[k],
		}
		out = append(out, j)
		if i >= 4 {
			break
		}
	}

	return out
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
