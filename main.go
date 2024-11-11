package main

import (
	"fmt"
	"strings"
)

type article struct {
	totalWords int
	words      []word
}
type vectorCompare struct {
	vectors []vector
}
type vector struct {
	allWords map[string]float32
}
type word struct {
	word           string
	numAppereances int
	docFreq        float32
}

func splitWords(document string) []string {
	document = strings.ToLower(document)
	words := strings.Split(document, " ")
	return words
}

func (vec *vector) articleGenerator(words []string) article {
	output := article{}
	bag := make(map[string]int)
	countWords := 0
	for i := range words {
		_, ok := bag[words[i]]
		if !ok {
			bag[words[i]] = 1
		} else {
			bag[words[i]]++
		}
		_, inVec := vec.allWords[words[i]]
		if !inVec {
			vec.allWords[words[i]] = 0
		}
		countWords++
	}
	output.totalWords = countWords
	bagOfWords := []word{}
	for i := range bag {
		df := float32(bag[i]) / float32(output.totalWords)
		bagOfWords = append(bagOfWords, word{
			word:           i,
			numAppereances: bag[i],
			docFreq:        df,
		})
	}
	output.words = bagOfWords
	return output
}

func idfVector(corpus []string) vectorCompare {
	var all []article
	emptyVect := vector{
		allWords: make(map[string]float32),
	}
	for i := range corpus {
		content := emptyVect.articleGenerator(splitWords(corpus[i]))
		all = append(all, content)
	}
	allVec := vectorCompare{
		vectors: make([]vector, len(all)),
	}
	for j := 0; j < len(all); j++ {
		tmp := emptyVect
		tmp.allWords = make(map[string]float32)
		for word := range emptyVect.allWords {
			tmp.allWords[word] = 0
		}
		for i, word := range all[j].words {
			tmp.allWords[word.word] = all[j].words[i].docFreq
		}
		allVec.vectors[j].allWords = tmp.allWords
	}
	return allVec
}

func main() {
	corpus := make([]string, 3)
	corpus[0] = "data science is one of the most important fields of science"
	corpus[1] = "Hello world"
	corpus[2] = "hello world"
	fmt.Printf("%v\n", idfVector(corpus))
}
