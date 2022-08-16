/*
👋 Welcome Gophers! In this challenge, you will be tasked with efficiently counting the word frequencies of a large body of text in Go!

You will have to implement a function which keeps track of the number of times a word appears in a body of text and then you will have to implement a further function which returns the top 5 most frequent words from highest to lowest.
*/
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
type ByScore []Word

func (ws ByScore) Len() int           { return len(ws) }
func (ws ByScore) Swap(x, y int)      { ws[x], ws[y] = ws[y], ws[x] }
func (ws ByScore) Less(x, y int) bool { return ws[x].Score > ws[y].Score }

func CountWords(text string) map[string]int {
	mapWs := make(map[string]int)
	xs := strings.Split(text, " ")

	for _, w := range xs {
		word := strings.Replace(w, ",", "", -1)
		word = strings.Replace(word, ".", "", -1)
		mapWs[word]++
	}

	return mapWs
}

func Top5Words(wordmap map[string]int) []Word {
	var Words []Word

	for k, v := range wordmap {
		w := Word{
			Value: k,
			Score: v,
		}
		Words = append(Words, w)
	}

	sort.Sort(ByScore(Words))
	return Words[:5]
}

func main() {
	fmt.Println("Word Frequency Test")

	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	results := CountWords(text)
	MostCommon := Top5Words(results)

	fmt.Println(MostCommon)
}
