package main

import (
	"strings"
)

func splitWords(document string) []string {
	document = strings.ToLower(document)
	words := strings.Split(document, " ")
	return removeStopWords(words)
}

func removeStopWords(input []string) []string {
	var output []string
	stopWords := []string{"i", "me", "my", "myself", "we", "our", "ours", "ourselves", "you", "you're", "you've", "you'll", "you'd",
		"your", "yours", "yourself", "yourselves", "he", "him", "his", "himself", "she", "she's", "her", "hers", "herself", "it", "it's", "its",
		"itself", "they", "them", "their", "theirs", "themselves", "what", "which", "who", "whom", "this", "that", "that'll", "these", "those", "am", "is",
		"are", "was", "were", "be", "been", "being", "have", "has", "had", "having", "do", "does", "did", "doing", "a", "an", "the", "and", "but", "if", "or",
		"because", "as", "until", "while", "of", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above",
		"below", "to", "from", "up", "down", "in", "out", "on", "off", "over", "under", "again", "further", "then", "once", "here", "there", "when", "where", "why",
		"how", "all", "any", "both", "each", "few", "more", "most", "other", "some", "such", "no", "nor", "not", "only", "own", "same", "so", "than", "too", "very", "s", "t", "can", "will", "just", "don",
		"don't", "should", "should've", "now", "d", "ll", "m", "o", "re", "ve", "y", "ain", "aren", "aren't", "couldn", "couldn't", "didn", "didn't", "doesn", "doesn't", "hadn",
		"hadn't", "hasn", "hasn't", "haven", "haven't", "isn", "isn't", "ma", "mightn", "mightn't", "mustn", "mustn't", "needn", "needn't",
		"shan", "shan't", "shouldn", "shouldn't", "wasn", "wasn't", "weren", "weren't", "won", "won't", "wouldn", "wouldn't"}
	for i := 0; i < len(input); i++ {
		add := true
		for j := 0; j < len(stopWords); j++ {
			if strings.ToLower(input[i]) == stopWords[j] {
				add = false
				break
			}
		}
		if add {
			output = append(output, strings.ToLower(input[i]))
		}
	}
	return output
}

func (vec *vector) articleGenerator(words []string) article {
	var indoc float64
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
		df := float64(bag[i]) / float64(output.totalWords)
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
