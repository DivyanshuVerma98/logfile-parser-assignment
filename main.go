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
	errorChan := make(chan string, 10)
	normalChan := make(chan string, 10)
	wg.Add(2)
	go func() {
		defer wg.Done()
		WriteReport(normalFileWriter, normalChan)
	}()
	go func() {
		defer wg.Done()
		WriteReport(errorFileWriter, errorChan)
	}()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		log := scanner.Text()
		wg.Add(1)
		go func(log string) {
			defer wg.Done()
			if strings.Contains(log, "ERROR") {
				errorChan <- log
			} else {
				normalChan <- log
			}
		}(log)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
	wg.Wait()
	close(errorChan)
	close(normalChan)
}

func WriteReport(writer io.Writer, ch <-chan string) {
	for value := range ch {
		_, err := writer.Write([]byte(value))
		if err != nil {
			log.Fatalln("Error writting into file: ", err)
			return
		}
	}
}

func main() {
	fmt.Println("Hello World")
	LogPraser("inputFile.log", "normalFile.log", "errorFile.log")
}
