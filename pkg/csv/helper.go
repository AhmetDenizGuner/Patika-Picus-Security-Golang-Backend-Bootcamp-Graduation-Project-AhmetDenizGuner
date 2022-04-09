package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

//ReadCsv reads csv that given path
func ReadCsv(filename string, startIndex int) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	//close the file
	defer f.Close()

	reader := csv.NewReader(f)

	//read line by line
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result [][]string

	for _, line := range lines[startIndex:] {
		result = append(result, line)
	}

	return result, nil
}

func ReadCSVWithWorkerPool(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	result := make([][]string, 0)
	numWps := 10
	jobs := make(chan []string, numWps)
	res := make(chan []string)

	var wg sync.WaitGroup
	worker := func(jobs <-chan []string, results chan<- []string) {
		for {
			select {
			case job, ok := <-jobs: // you must check for readable state of the channel.
				if !ok {
					return
				}
				results <- job
			}
		}
	}

	// init workers
	for w := 0; w < numWps; w++ {
		wg.Add(1)
		go func() {
			// this line will exec when chan `res` processed output at line 107 (func worker: line 71)
			defer wg.Done()
			worker(jobs, res)
		}()
	}

	go func() {
		for {
			rStr, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				break
			}
			jobs <- rStr
		}
		close(jobs) // close jobs to signal workers that no more job are incoming.
	}()

	go func() {
		wg.Wait()
		close(res) // when you close(res) it breaks the below loop.
	}()

	for r := range res {
		result = append(result, r)
	}

	return result, nil
}
