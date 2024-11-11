package main

import (
	"fmt"
	"math"
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
	termFreq       float32
}

func splitWords(document string) []string {
	document = strings.ToLower(document)
	words := strings.Split(document, " ")
	return words
}

func (vec *vector) articleGenerator(words []string) article {
	var indoc float32
	indoc = 0
	output := article{}
	bag := make(map[string]int)
	countWords := 0
	for i := range words {
		_, ok := bag[words[i]]
		if !ok {
			bag[words[i]] = 1
			indoc = 1
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
			termFreq:       df,
			docFreq:        indoc,
		})
	}
	output.words = bagOfWords
	return output
}

func tfVector(corpus []string) vectorCompare {
	var all []article
	count := make([]int, len(corpus))
	emptyVect := vector{
		allWords: make(map[string]float32),
	}
	for i := range corpus {
		content := emptyVect.articleGenerator(splitWords(corpus[i]))
		count[i] = content.totalWords
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
			tmp.allWords[word.word] = all[j].words[i].termFreq / float32(count[j])
		}
		allVec.vectors[j].allWords = tmp.allWords
	}
	return allVec
}

func idfVector(corpus []string) vector {
	var all []article
	emptyVect := vector{
		allWords: make(map[string]float32),
	}
	for i := range corpus {
		content := emptyVect.articleGenerator(splitWords(corpus[i]))
		for j := range content.words {
			content.words[j].termFreq = 0
		}
		all = append(all, content)
	}
	tmp := emptyVect
	tmp.allWords = make(map[string]float32)
	for word := range emptyVect.allWords {
		tmp.allWords[word] += 0
	}
	for j := 0; j < len(all); j++ {
		for i, word := range all[j].words {
			tmp.allWords[word.word] += all[j].words[i].docFreq
		}
	}
	for i := range tmp.allWords {
		tmp.allWords[i] = float32(math.Log10(float64(len(all)) / float64(tmp.allWords[i])))
	}
	return tmp
}

func tfIdfVec(tf vectorCompare, idf vector) []vector {
	res := make([]vector, len(tf.vectors))
	for i := 0; i < len(tf.vectors); i++ {
		for word := range tf.vectors[i] {
			res[i] = tf.vectors[item] * idf[item]
		}
	}
	return make([]vector, 1)
}

func main() {
	corpus := make([]string, 3)
	corpus[0] = "data science is one of the most important fields of science"
	corpus[1] = "Hello world"
	corpus[2] = "hello hello world"
	tfVector := tfVector(corpus)
	fmt.Printf("%v\n", tfVector)
	idfVector := idfVector(corpus)
	fmt.Printf("%v\n", idfVector)
	fmt.Printf("%v\n", tfIdfVec(tfVector, idfVector))
}
