package main

import (
	"errors"
	"math"
)

func dotProduct(vector1 vector, vector2 vector) (float64, error) {
	if len(vector1.allWords) != len(vector2.allWords) {
		return 0.0, errors.New("can not create dot product of two different length vectors")
	}
	var dotProduct float64
	dotProduct = 0
	for k := range vector1.allWords {
		dotProduct += vector1.allWords[k] * vector2.allWords[k]
	}
	return dotProduct, nil
}

func norm(vector1 vector) float64 {
	var sumSquares float64
	sumSquares = 0
	for _, v := range vector1.allWords {
		sumSquares += float64(v * v)
	}
	return math.Sqrt(sumSquares)
}

func cosineSimiliarity(vector1, vector2 vector) (float64, error) {
	dot, err := dotProduct(vector1, vector2)
	if err != nil {
		return 0.0, err
	}
	cosine := dot / (norm(vector1) * norm(vector2))
	return cosine, nil
}
