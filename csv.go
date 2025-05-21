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

func processCsvFile(records *[][]string) map[string]record {
	var recordsMap = make(map[string]record)
	rec := *records
	for i := 0; i < len(rec); i++ {
		var r = record{
			trackId: rec[i][0],
			albumId: rec[i][1],
			title:   rec[i][2],
		}
		recordsMap[r.trackId] = r
	}
	return recordsMap
}

type record struct {
	trackId string
	albumId string
	title   string
}

func main() {
	records := readCsvFile("data/tracks.csv")
	var allRecordsMap = processCsvFile(&records)
	for _, rec := range allRecordsMap {
		fmt.Println("Record: ", rec)
	}
}
