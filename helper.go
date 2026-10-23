package main

import (
	"fmt"
	"os"
	"strings"
)

func writeFile(filename string, s string) error {
	//Converts the string to a byte for writing to the file
	data := []byte(s)
	// Writes data to the files (0666: read and write permission for everyone)
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		// Error handling (e.g, file creation issues)
		fmt.Printf("Something went wrong !! %v", err)
		os.Exit(1) // Exit with error code
	}
	return err
}

func readFromFile(filename string) string {
	// Reads file contents
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return string(data)
}

func countCharacters(text string) int {
	//Length of the string represents character count
	return len(text)
}

func Wordcount(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countLines(text string) int {
	lines := strings.Split(text, "\n")
	return len(lines)
}

func countFrequency(s []string) map[string]int {
	//Stores word frequences
	m := make(map[string]int)
	for _, word := range s {
		count, ok := m[word]
		if !ok {
			count = 1
			//If word is not found initizalize its count to 1
			// if it is found increment its count
		} else {
			count++
		}
		m[word] = count
	}
	return m
}

func printingFrequency(m map[string]int, s []string) {
	//keeps track of visited words (used to avoid duplicates (e.g the the))
	visited := make(map[string]bool)
	fmt.Println("Printing frequency of each word : ")
	for _, word := range s {
		if !visited[word] {
			fmt.Printf("%s => %d time(s), ", word, m[word])
			visited[word] = true
		}
	}
	fmt.Println()
}

func countSentences(text string) int {
	count := 0

	for _, ch := range text {

		if ch == '.' || ch == '!' || ch == '?' {
			count++

		}
	}
	return count
}

func averageWordLength(text string) float64 {
	words := strings.Fields(text)

	if len(words) == 0 {
		return 0
	}

	totalLength := 0
	validWords := 0

	for _, word := range words {
		cleaned := strings.Trim(word, ".,!?;:\"()[]")

		if cleaned != "" {
			totalLength += len(cleaned)
			validWords++
		}
	}

	if validWords == 0 {
		return 0
	}

	return float64(totalLength) / float64(validWords)
}
