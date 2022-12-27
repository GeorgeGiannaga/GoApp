package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) []string {
	words := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func main() {
	var word string
	filename := "dat.txt"
	fmt.Println("Please specify the word , that you want to be counted in the text of the file")
	fmt.Scan(&word)
	text := ReadFile(filename)
	var j int
	for i := 0; i < len(text); i++ {
		if text[i] == string(word) {
			j++
		}
	}
	if j == 0 {
		fmt.Printf("There is no such word found in the text")
	} else {
		fmt.Printf("'%s' word has been identified in the next %v time/s", word, j)
	}
}
