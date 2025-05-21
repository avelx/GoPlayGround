package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func processCsvFile(records *[][]string) map[int]record {
	var recordsMap = make(map[int]record)
	rec := *records
	for i := 0; i < len(rec); i++ {
		id, _ := strconv.Atoi(rec[i][0])
		var r = record{
			trackId: id,
			albumId: rec[i][1],
			title:   rec[i][2],
		}
		recordsMap[r.trackId] = r
	}
	return recordsMap
}

type record struct {
	trackId int
	albumId string
	title   string
}

func main() {
	records := readCsvFile("data/tracks.csv")
	var allRecordsMap = processCsvFile(&records)

	for _, rec := range allRecordsMap {
		//filter our all records: between [100-200]
		//if rec.trackId >= 100 && rec.trackId <= 200 {
		//	fmt.Println("Record: ", rec)
		//}
		matched, err := regexp.Match(`Two.*`, []byte(rec.title))
		if err != nil {
			log.Fatal(err)
		}
		if matched {
			fmt.Println("Record: ", rec)
		}
	}
}
