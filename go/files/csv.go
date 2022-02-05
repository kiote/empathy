package files

import (
	"os"
	"log"
	"encoding/csv"
	"fmt"
)

func CreateCSV(name string, dataDir string) (*os.File, error) {
	os.Mkdir(dataDir, 0775)
	csvFile, err := os.Create(dataDir + "/" + name + ".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	return csvFile, err
}

func WriteCSV(csvFile *os.File, data [][]string) {
	csvwriter := csv.NewWriter(csvFile)
	fmt.Printf("%v", data)
	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}