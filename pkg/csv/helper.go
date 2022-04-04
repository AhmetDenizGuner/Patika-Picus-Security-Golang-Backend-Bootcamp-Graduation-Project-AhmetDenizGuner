package csv

import (
	"encoding/csv"
	"os"
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
