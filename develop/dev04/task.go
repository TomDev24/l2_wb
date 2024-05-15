package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words []string) map[string][]string {
	anagramSets := make(map[string][]string)
	result := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(strings.ToLower(word))
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	for _, value := range anagramSets {
		if len(value) > 1 {
			sort.Strings(value)
			result[value[0]] = value
		}
	}

	return result
}

func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagramSets(words)

	for key, value := range anagramSets {
		fmt.Println(key, value)
	}
}
