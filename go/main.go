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
			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", req.PostForm)
			name := req.FormValue("name")
			address := req.FormValue("address")
			fmt.Fprintf(w, "Name = %s\n", name)
			fmt.Fprintf(w, "Address = %s\n", address)
			saveDemographic(w, req)
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
	csvFile.Close()
	return csvFile, err
}

func writeCSV(file *os.File, data [][]string) {
	w := csv.NewWriter(file)
	fmt.Printf("%v", data)
	defer w.Flush()
	w.WriteAll(data)
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/demographic", demgoraphic)

	fmt.Printf("Listening 8090\n")
    http.ListenAndServe(":8090", nil)
}