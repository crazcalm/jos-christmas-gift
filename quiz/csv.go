package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

//ReadCSV -- Parses csv file
func ReadCSV(path string) (records [][]string, err error) {
	//Make sure the file exists
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = fmt.Errorf("file: %s does not exist", path)
		return
	}

	//Open the file
	file, err := os.Open(path)
	if err != nil {
		return
	}

	//Read the file
	r := csv.NewReader(file)
	r.LazyQuotes = true // Needed to except the existence of quotes within the statement fields

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
		records = append(records, record)
	}

	return
}
