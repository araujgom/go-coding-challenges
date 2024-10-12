package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func countLines(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func countWords(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

func countBytes(r io.Reader) (int64, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return int64(len(content)), nil
}

func countCharacters(content []byte) int {
	return utf8.RuneCountInString(string(content))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s [options] <filename>\n", os.Args[0])
	}

	// If no filename is provided, default to stdin
	var input io.Reader
	var filename string
	if len(os.Args) == 2 {
		input = os.Stdin
	} else {
		filename = os.Args[len(os.Args)-1]
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Error opening file: %v\n", err)
		}
		defer file.Close()
		input = file
	}

	var printAll bool
	var printBytes, printLines, printWords, printChars bool

	// Parse options
	if len(os.Args) == 2 {
		// No option provided, print all
		printAll = true
	} else {
		for _, arg := range os.Args[1 : len(os.Args)-1] {
			switch arg {
			case "-c":
				printBytes = true
			case "-l":
				printLines = true
			case "-w":
				printWords = true
			case "-m":
				printChars = true
			default:
				log.Fatalf("Invalid option: %s\n", arg)
			}
		}
	}

	var bytes int64
	var lines, words, chars int
	content, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}

	// Get byte count
	if printBytes || printAll {
		bytes = int64(len(content))
	}

	// Get line count
	if printLines || printAll {
		lines = countLines(bufio.NewReader(input))
	}

	// Get word count
	if printWords || printAll {
		words = countWords(bufio.NewReader(input))
	}

	// Get character count
	if printChars || printAll {
		chars = countCharacters(content)
	}

	// Print results based on selected options
	if printAll {
		fmt.Printf("%d %d %d %d %s\n", lines, words, bytes, chars, filename)
	} else {
		if printLines {
			fmt.Printf("%d ", lines)
		}
		if printWords {
			fmt.Printf("%d ", words)
		}
		if printBytes {
			fmt.Printf("%d ", bytes)
		}
		if printChars {
			fmt.Printf("%d ", chars)
		}
		fmt.Printf("%s\n", filename)
	}
}
