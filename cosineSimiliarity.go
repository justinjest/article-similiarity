package main

import (
	"errors"
	"fmt"
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

func mostSimiliar(vector1 vector, searchVector []vector) vector {
	similiarityRanking := make([]float64, len(searchVector))
	for i := 0; i < len(similiarityRanking); i++ {
		sim, err := cosineSimiliarity(vector1, searchVector[i])
		if err != nil {
			sim = 0
			fmt.Print("Invalid comparison due to different sized arrays")
		} else {
			similiarityRanking[i] = sim
		}
	}
	ans := float64(0)
	loc := 0
	for i := range similiarityRanking {
		if similiarityRanking[i] > ans {
			ans = similiarityRanking[i]
			loc = i
		}
	}
	return searchVector[loc]
}
