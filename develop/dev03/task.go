package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"sort"
	"strconv"
	"strings"
)

func reverseLines(lines []string) {
	for i := 0; i < len(lines)/2; i++ {
		j := len(lines) - i - 1
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func removeDuplicates(lines []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, line := range lines {
		if !encountered[line] {
			encountered[line] = true
			result = append(result, line)
		}
	}
	return result
}

func sortFile(filePath string, column int, numeric, reverse, unique bool) []string {
	if filePath == "" {
		log.Fatal("Не указан путь к файлу")
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fileContent), "\n")
	sort.SliceStable(lines, func(i, j int) bool {
		lineA := lines[i]
		lineB := lines[j]

		if column > 0 && column <= len(strings.Fields(lineA)) && column <= len(strings.Fields(lineB)) {
			fieldA := strings.Fields(lineA)[column-1]
			fieldB := strings.Fields(lineB)[column-1]

			if numeric {
				numA, errA := strconv.Atoi(fieldA)
				numB, errB := strconv.Atoi(fieldB)

				if errA == nil && errB == nil {
					return numA < numB
				}
			}

			return fieldA < fieldB
		}

		return lineA < lineB
	})

	if reverse {
		reverseLines(lines)
	}

	if unique {
		lines = removeDuplicates(lines)
	}

	return lines
}

func main() {
	filePath := flag.String("file", "", "путь к файлу для сортировки")
	column := flag.Int("k", 0, "номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()
	lines := sortFile(*filePath, *column, *numeric, *reverse, *unique)

	output := strings.TrimSpace(strings.Join(lines, "\n"))
	outputFilePath := *filePath + ".sorted"
	err := os.WriteFile(outputFilePath, []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Отсортированные данные сохранены в файл: %s\n", outputFilePath)
}
