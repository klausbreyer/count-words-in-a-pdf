package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/dslipak/pdf"
)

type WordCount struct {
	Word  string
	Count int
}

func countWords(str string) []WordCount {
	wordList := strings.Fields(str)
	wordCountMap := make(map[string]int)

	for _, word := range wordList {
		if len(word) < 5 {
			continue
		}
		wordCountMap[word]++
	}

	// Convert map to slice for sorting
	var wordCounts []WordCount
	for word, count := range wordCountMap {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}

	// Sort slice in descending order of count, and alphabetically for same count
	sort.SliceStable(wordCounts, func(i, j int) bool {
		if wordCounts[i].Count != wordCounts[j].Count {
			return wordCounts[i].Count > wordCounts[j].Count
		}
		return wordCounts[i].Word < wordCounts[j].Word // Alphabetical order for same frequency
	})

	return wordCounts
}

func readPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	// remember close file
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Bitte geben Sie den Pfad zur PDF-Datei an.\n")
		os.Exit(1)
	}

	pdfPath := os.Args[1]
	text, err := readPdf(pdfPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim readpdf: %v\n", err)
		os.Exit(1)
	}

	// Nur Wörter extrahieren und zählen
	reg, err := regexp.Compile("[^a-zA-Z0-9äöüÄÖÜß]+")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Erstellen des Regex: %v\n", err)
		os.Exit(1)
	}

	processedString := reg.ReplaceAllString(text, " ")
	wordCounts := countWords(processedString)

	// Top 100 Wörter ausgeben
	for i, wc := range wordCounts {
		if i == 100 {
			break
		}
		fmt.Printf("Wort: %s, Anzahl: %d\n", wc.Word, wc.Count)
	}
}
