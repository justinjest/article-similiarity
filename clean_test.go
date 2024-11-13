package main

import (
	"log"
	"testing"
)

func TestCleanWords(t *testing.T) {
	expected := []string{"hello", "world"}
	result := removeStopWords([]string{"Hello", "world"})
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] || len(expected) != len(result) {
			log.Fatalf("error with cleaning words expected %v, got %v\n", expected, result)
		}
	}
	expected = []string{"hello", "world"}
	result = removeStopWords([]string{"Hello", "world", "i"})
	for i := 0; i < len(expected); i++ {
		if expected[i] != result[i] || len(expected) != len(result) {
			log.Fatalf("error with cleaning words expected %v, got %v\n", expected, result)
		}
	}
}
