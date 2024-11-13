package main

import (
	"fmt"
)

func main() {
	corpus := make([]string, 5)
	corpus[0] = "data science is one of the most important fields of science"
	corpus[1] = "Hello world"
	corpus[2] = "hello hello world"
	corpus[3] = "hello Hello world i i i i i"
	corpus[4] = "You shouldn't end the world"
	tfVector := tfVector(corpus)
	idfVector := idfVector(corpus)
	tfIdf := tfIdfVec(tfVector, idfVector)
	res := mostSimiliar(tfIdf[2], tfIdf[0:2])
	fmt.Printf("most similar is %v\n", res)
}
