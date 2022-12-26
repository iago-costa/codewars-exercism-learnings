package main

import (
	"fmt"
	"example.com/reverse_words"
)

func main() {
	fmt.Println("Hello World")
	result := reverse_words.ReverseWords("The quick brown fox jumps over the lazy dog.")
	fmt.Println(result)
}

