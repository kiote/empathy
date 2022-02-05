package main

import (
    "fmt"
	"net/http"
	"log"
	"time"
	"os"
	"strconv"
	"encoding/csv"
	"example.com/m/v2/sensors"
)

var currentTimestamp = time.Now().Unix()
var dataDir = "../" +  fmt.Sprintf("%d", currentTimestamp)

//
// HTTP handlers
//

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

// Handle "/eq" requests
func eq(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":		
			 http.ServeFile(w, req, "../static/eq.html")
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			saveEq(w, req)
			http.Redirect(w, req, "/video", 302)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
}

// Handle "/video" requests
func video(w http.ResponseWriter, req *http.Request) {
	number := 1
	if currentTimestamp % 2 == 0 {
		number = 2
	}

	switch req.Method {
		case "GET":		
			 http.ServeFile(w, req, "../static/video" + strconv.Itoa(number) +".html")
			 sensors.StartExperiment()
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
		}
}

// Handle "/se1" requests
func se1(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":		
			 http.ServeFile(w, req, "../static/se1.html")
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			saveSe(1, w, req)
			if (seFileExists(2)) {
				http.Redirect(w, req, "/done", 302)
			} else {
				http.ServeFile(w, req, "../static/video2.html")
			}
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
		}
}

// Handle "/se2" requests
func se2(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "GET":		
			 http.ServeFile(w, req, "../static/se2.html")
		case "POST":
			if err := req.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			saveSe(2, w, req)
			if (seFileExists(1)) {
				http.Redirect(w, req, "/done", 302)
			} else {
				http.ServeFile(w, req, "../static/video1.html")
			}
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported.")
		}
}

func done(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../static/done.html")
}

//
// End of HTTP handlers
//


//
// Data saving functions 
//

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

//Save EQ test data
func saveEq(w http.ResponseWriter, req *http.Request) {
	eqData := [][]string{
		{"q1", "q2", "q3", "q4", "q5", "q6", "q7", "q8", "q9", "q10",
		"q11", "q12", "q13", "q14", "q15", "q16", "q17", "q18", "q19", "q20",
		"q21", "q22", "q23", "q24", "q25", "q26", "q27", "q28", "q29", "q30",
		"q31", "q32", "q33", "q34", "q35", "q36", "q37", "q38", "q39", "q40"},
		{ req.FormValue("q1"), req.FormValue("q2"), req.FormValue("q3"), req.FormValue("q4"), req.FormValue("q5"), 
		req.FormValue("q6"), req.FormValue("q7"), req.FormValue("q8"), req.FormValue("q9"), req.FormValue("q10"), 
		req.FormValue("q11"), req.FormValue("q12"), req.FormValue("q13"), req.FormValue("q14"), req.FormValue("q15"), 
		req.FormValue("q16"), req.FormValue("q17"), req.FormValue("q18"), req.FormValue("q19"), req.FormValue("q20"),
		req.FormValue("q21"), req.FormValue("q22"), req.FormValue("q23"), req.FormValue("q24"), req.FormValue("q25"), 
		req.FormValue("q26"), req.FormValue("q27"), req.FormValue("q28"), req.FormValue("q29"), req.FormValue("q30"),
		req.FormValue("q31"), req.FormValue("q32"), req.FormValue("q33"), req.FormValue("q34"), req.FormValue("q35"), 
		req.FormValue("q36"), req.FormValue("q37"), req.FormValue("q38"), req.FormValue("q39"), req.FormValue("q40")},
	}

	file, _ := createCSV("eq")
	writeCSV(file, eqData)
}

func saveSe(num int, w http.ResponseWriter, req *http.Request) {
	seData := [][]string{
		{"q1", "q2", "q3", "q4", "q5"},
		{ req.FormValue("q1"), req.FormValue("q2"), req.FormValue("q3"), req.FormValue("q4"), req.FormValue("q5")},
	}

	file, _ := createCSV("se" + strconv.Itoa(num))
	writeCSV(file, seData)
}

//
// End of Data saving functions
//

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

func seFileExists(num int) bool {
	exists := false
	fmt.Printf(dataDir + "/se" + strconv.Itoa(num) + ".csv")
	if _, err := os.Stat(dataDir + "/se" + strconv.Itoa(num) + ".csv"); err == nil {
		fmt.Printf("File exists")
		exists = true
	}

	return exists
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/demographic", demgoraphic)
	http.HandleFunc("/eq", eq)
	http.HandleFunc("/video", video)
	http.HandleFunc("/se1", se1)
	http.HandleFunc("/se2", se2)
	http.HandleFunc("/done", done)

	fmt.Printf("Listening 8090\n")
    http.ListenAndServe(":8090", nil)
}