package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var ms string

	for i := 0; i < (len(word1) + len(word2)); i++ {

		if i < len(word1) {
			ms += string(word1[i])
		}

		if i < len(word2) {
			ms += string(word2[i])
		}

	}

	return ms
}

func main() {
	fmt.Println("result:", mergeAlternately("abc", "pqr"))
	fmt.Println("result:", mergeAlternately("ab", "pqrs"))
	fmt.Println("result:", mergeAlternately("abcd", "pq"))
}
