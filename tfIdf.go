package main

import (
	"math"
	"sort"
)

type article struct {
	totalWords int
	words      []word
}
type vectorCompare struct {
	vectors []vector
}
type vector struct {
	allWords map[string]float64
}
type word struct {
	word           string
	numAppereances int
	docFreq        float64
	termFreq       float64
}

func tfVector(corpus []string) vectorCompare {
	var all []article
	count := make([]int, len(corpus))
	emptyVect := vector{
		allWords: make(map[string]float64),
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
		tmp.allWords = make(map[string]float64)
		for word := range emptyVect.allWords {
			tmp.allWords[word] = 0
		}
		for i, word := range all[j].words {
			tmp.allWords[word.word] = all[j].words[i].termFreq / float64(count[j])
		}

		keys := make([]string, 0, len(tmp.allWords))

		for k := range tmp.allWords {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		allVec.vectors[j].allWords = tmp.allWords
		for _, k := range keys {
			allVec.vectors[j].allWords[k] = tmp.allWords[k]
		}
	}
	return allVec
}

func idfVector(corpus []string) vector {
	var all []article
	emptyVect := vector{
		allWords: make(map[string]float64),
	}
	for i := range corpus {
		content := emptyVect.articleGenerator(splitWords(corpus[i]))
		for j := range content.words {
			content.words[j].termFreq = 0
		}
		all = append(all, content)
	}
	tmp := emptyVect
	tmp.allWords = make(map[string]float64)
	for word := range emptyVect.allWords {
		tmp.allWords[word] = 0
	}
	for j := 0; j < len(all); j++ {
		for i, word := range all[j].words {
			tmp.allWords[word.word] += all[j].words[i].docFreq
		}
	}
	for i := range tmp.allWords {
		tmp.allWords[i] = math.Log10(float64(len(all)) / float64(tmp.allWords[i]))
	}
	return tmp
}

func tfIdfVec(tf vectorCompare, idf vector) []vector {
	res := make([]vector, len(tf.vectors))
	for i := 0; i < len(tf.vectors); i++ {
		res[i].allWords = make(map[string]float64, len(idf.allWords))
		for word := range tf.vectors[i].allWords {
			tmp := tf.vectors[i].allWords[word] * idf.allWords[word]
			res[i].allWords[word] = tmp
		}
	}
	return res
}
