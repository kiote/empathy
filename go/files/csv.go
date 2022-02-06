package files

import (
	"os"
	"log"
	"encoding/csv"
	"fmt"
	"example.com/m/v2/general"
)

func CreateCSV(name string) (*os.File, error) {
	os.Mkdir(settings.DataDir, 0775)
	csvFile, err := os.Create(settings.DataDir + "/" + name + ".csv")

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