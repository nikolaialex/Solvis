package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const shortForm = "02.01.06 15:04:05"
const inFolder = "in/*.txt"
const outFolder = "out"

func main() {
	wg := new(sync.WaitGroup)

	inputFiles, err := filepath.Glob(inFolder)
	if err != nil {
		panic(err)
	}

	// Process each file in a go routine.
	for _, filePath := range inputFiles {
		wg.Add(1)
		go processFile(filePath, wg)
	}

	wg.Wait()
}

func processFile(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing file: " + filePath)
	csvFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'

	csvLines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	exportFilePath := filepath.Join(outFolder, filepath.Base(filePath))
	f, err := os.Create(exportFilePath)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, line := range csvLines {
		var newLine []string
		timestamp, err := time.Parse(shortForm, line[0]+" "+line[1])
		newLine = append(newLine, strconv.FormatInt(timestamp.Unix(), 10))
		if err != nil {
			panic(err)
		}
		for index, value := range line[2:] {
			convertedValue, _ := strconv.ParseFloat(value, 64)
			if index < 18 {
				convertedValue = convertedValue * 0.1
			}
			newLine = append(newLine, fmt.Sprintf("%.2f", convertedValue))
		}
		f.WriteString(strings.Join(newLine, ",") + "\n")
	}
}
