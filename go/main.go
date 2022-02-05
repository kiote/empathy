package main

import (
    "fmt"
	"net/http"
	"log"
	"time"
	"os"
	"strconv"
	"encoding/csv"
)

var currentTimestamp = time.Now().Unix()
var dataDir = "../" +  fmt.Sprintf("%d", currentTimestamp)

func index(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../static/index.html")
}

// Handle "/demographic" requests
func demgoraphic(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":		
			 http.ServeFile(w, req, "../static/demography.html")
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			saveDemographic(w, req)
			http.Redirect(w, req, "/eq", 302)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
}

//Save demographic data
func saveDemographic(w http.ResponseWriter, req *http.Request) {
	drugs := req.FormValue("nodrugs")
	nodrugs := drugs != "" 
	age := req.FormValue("age")
	race := req.FormValue("race")
	sex := req.FormValue("sex")

	demographicData := [][]string{
		{"Age", "Race", "Sex", "No drugs"},
		{ age, race, sex, strconv.FormatBool(nodrugs) },
	}

	file, _ := createCSV("demographic")
	writeCSV(file, demographicData)
}

func createCSV(name string) (*os.File, error) {
	os.Mkdir(dataDir, 0775)
	csvFile, err := os.Create(dataDir + "/" + name + ".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	return csvFile, err
}

func writeCSV(csvFile *os.File, data [][]string) {
	csvwriter := csv.NewWriter(csvFile)
	fmt.Printf("%v", data)
	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/demographic", demgoraphic)

	fmt.Printf("Listening 8090\n")
    http.ListenAndServe(":8090", nil)
}