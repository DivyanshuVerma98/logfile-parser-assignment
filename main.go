package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// LogPraser processes the input log file and
// categorizes logs into normal and error files
func LogPraser(inputFileName string, normalFileName string, errorFileName string) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalln("Error while opening input file: ", err)
	}
	defer inputFile.Close()
	normalFileWriter, err := os.Create(normalFileName)
	if err != nil {
		log.Fatalln("Error while creating normal file: ", err)
	}
	defer normalFileWriter.Close()
	errorFileWriter, err := os.Create(errorFileName)
	if err != nil {
		log.Fatalln("Error while creating error file: ", err)
	}

	defer errorFileWriter.Close()
	var wg sync.WaitGroup
	var logWg sync.WaitGroup

	// Buffered channels for normal and error logs
	errorChan := make(chan string, 10)
	normalChan := make(chan string, 10)

	// Start goroutines to write to files
	wg.Add(2)
	go func() {
		defer wg.Done()
		WriteReport(normalFileWriter, normalChan)
	}()
	go func() {
		defer wg.Done()
		WriteReport(errorFileWriter, errorChan)
	}()

	// Scan the input file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		log := scanner.Text()
		logWg.Add(1)
		go func(log string) {
			defer logWg.Done()
			logList := strings.Split(log, " ")
			if logList[2] == "ERROR" {
				errorChan <- log
			} else {
				normalChan <- log
			}
		}(log)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	// Wait for all log processing goroutines to finish
	logWg.Wait()
	close(errorChan)
	close(normalChan)
	wg.Wait()
}

// WriteReport writes logs from the channel to the provided writer
func WriteReport(writer io.Writer, ch <-chan string) {
	for value := range ch {
		_, err := writer.Write([]byte(value + "\n"))
		if err != nil {
			log.Fatalln("Error writting into file: ", err)
			return
		}
	}
}

func main() {
	fmt.Println("Stating the program ... 🚀")
	LogPraser("inputFile.log", "normalFile.log", "errorFile.log")
	fmt.Println("Done 🔥")
}
