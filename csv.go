package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

/*
Read a simple CSV file and raise an error if any issues
*/
func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

type record struct {
	trackId string
	albumId string
	title   string
}

func main() {
	records := readCsvFile("data/tracks.csv")

	var recordsMap = make(map[string]record)

	for i := 0; i < len(records); i++ {
		var record = record{
			trackId: records[i][0],
			albumId: records[i][1],
			title:   records[i][2],
		}
		recordsMap[record.trackId] = record
		//fmt.Println(record)
		//fmt.Println(records[i][0])
	}

	fmt.Println(recordsMap)
}
