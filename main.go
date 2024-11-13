package main

import (
	"fmt"
)

func main() {
	corpus := make([]string, 3)
	corpus[0] = "data science is one of the most important fields of science"
	corpus[1] = "Hello world"
	corpus[2] = "hello hello world"
	tfVector := tfVector(corpus)
	idfVector := idfVector(corpus)
	tfIdf := tfIdfVec(tfVector, idfVector)
	res := mostSimiliar(tfIdf[2], tfIdf[0:2])
	fmt.Printf("most similar is %v\n", res)
}
