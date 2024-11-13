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
	res, err := cosineSimiliarity(tfIdf[0], tfIdf[1])
	fmt.Printf("%v\n", tfIdf[0])
	if err != nil {
		fmt.Printf("Error recevied %v\n", err)
		return
	}
	fmt.Printf("cosine is %v\n", res)
}
