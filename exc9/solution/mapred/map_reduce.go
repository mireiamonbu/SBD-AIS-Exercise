package mapred

import (
	"regexp"
	"strings"
	"sync"
)

type MapReduce struct {
}

// todo implement mapreduce
func (mr *MapReduce) wordCountMapper(text string) []KeyValue {
	// First clean the text
	re := regexp.MustCompile("[^a-zA-Z]+")
	text = re.ReplaceAllLiteralString(text, " ")

	// Turn the text into minusc
	text = strings.ToLower(text)

	// Split the text into words
	words := strings.Fields(text)

	// Now with every word, create a slice and add freq 1
	var result []KeyValue
	for _, word := range words {
		result = append(result, KeyValue{Key: word, Value: 1})
	}
	return result
}

func (mr *MapReduce) wordCountReducer(key string, values []int) KeyValue {
	// Sum the frequencies of the word
	var sum int
	for _, value := range values {
		sum += value
	}

	// return the word and its total frquency
	return KeyValue{Key: key, Value: sum}
}

// principal function
func (mr *MapReduce) Run(input []string) map[string]int {
	// Wait until all the gorutines have finished
	var wg sync.WaitGroup

	// Map the result, grouped by word
	intermediate := make(map[string][]int)

	// create a chanel to recolect all the results
	resultsChannel := make(chan KeyValue)

	// Map
	for _, line := range input {
		wg.Add(1)

		go func(line string) {
			defer wg.Done()

			mapped := mr.wordCountMapper(line)

			for _, keyval := range mapped {
				resultsChannel <- keyval
			}
		}(line)
	}

	// create a goroutine to wait every goroutine of the map to be finished
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	// get all the result from the chane and group by word
	for keyval := range resultsChannel {
		intermediate[keyval.Key] = append(intermediate[keyval.Key], keyval.Value)
	}

	// reduce words frequency
	finalResults := make(map[string]int)
	for key, values := range intermediate {
		reduced := mr.wordCountReducer(key, values)
		finalResults[reduced.Key] = reduced.Value
	}

	return finalResults
}
