package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "Textfile.txt"

	_, err := os.Stat(filename)
	if err == nil {
		os.Remove(filename)
	}

	writeFile("Textfile.txt", strings.Join(os.Args[1:], " "))

	text := readFromFile("Textfile.txt")
	s := strings.Fields(text)

	characterCount := countCharacters(text)
	wordCount := Wordcount(text)
	lineCount := countLines(text)
	sentenceCount := countSentences(text)
	avgWordlength := averageWordLength(text)

	m := countFrequency(s)

	printingFrequency(m, s)

	fmt.Printf("\nCharacter count: %d\n", characterCount)
	fmt.Printf("Word count: %d\n", wordCount)
	fmt.Printf("Line count: %d\n", lineCount)
	fmt.Printf("Sentence count: %d\n", sentenceCount)
	fmt.Printf("Average word length: %.2f\n", avgWordlength)
}
