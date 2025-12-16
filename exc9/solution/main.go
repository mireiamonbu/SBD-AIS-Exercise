package main

import (
	"bufio"
	"exc9/mapred"
	"fmt"
	"log"
	"os"
)

// Main function
func main() {
	// todo read file
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read line by line of the file
	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// verify if any error happends during the lecture
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	for word, count := range results {
		fmt.Printf("%s: %d\n", word, count)
	}
}
