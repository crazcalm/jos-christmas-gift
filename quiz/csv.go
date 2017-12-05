package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//ReadCSV -- Parses csv file
func ReadCSV() [][]string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
	for _, f := range files {
		fmt.Println(f.Name())
	}
	testFile := "testing.csv"
	_, err = os.Stat(testFile)
	if os.IsNotExist(err) {
		log.Fatalf("file: %s does not exist", testFile)
	}

	file, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	var records [][]string
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

	return records
}
