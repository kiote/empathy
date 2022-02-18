package sensors

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"encoding/json"
	"example.com/m/v2/files"
	"net"
    "os"
)

type Samples struct {
	GetExperimentSamples [][]interface{} `json:"GetExperimentSamples"`
}

var (
	serverUrl = "http://localhost"
	serverPort = "22002"
	serverPath = "NeuLogAPI"
	
	startExperiment = "StartExperiment:[GSR],[1],[9],[360]"
	stopExperiment = "StopExperiment"
	getSamples = "GetExperimentSamples:[GSR],[1]"

	tcpServerAddress = "127.0.0.1"
	tcpServerPort = "8089"
	startStimuliTCPMessage = "M;2;;;Started;Stimuli started;D;\r\n"
	stopStimuliTCPMessage = "M;2;;;Finished;Stimuli finished;D;\r\n"
)

func StartExperiment() {
	fmt.Printf("GSR Sensor initializing...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + startExperiment + "\n")
	httpCall(startExperiment)
	tcpCall(startStimuliTCPMessage)
}

func StopExperiment() {
	fmt.Printf("GSR Data collection stop...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + stopExperiment + "\n")
	httpCall(stopExperiment)
	tcpCall(stopStimuliTCPMessage)
}

func SaveExperiment(fileNum int) {
	fmt.Printf("GSR Data collection save...\n")
	fmt.Printf("To file %d", fileNum)
	smaples, _ := httpCall(getSamples)
	saveSamples(fileNum, smaples)
}

func httpCall(command string) (Samples, error) {
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + command + "\n")
	var result Samples
	resp, err := http.Get(serverUrl + ":" + serverPort + "/" + serverPath + "?" + command)
	if err != nil {
		fmt.Printf("%e \n", err)
		return result, err
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body) // response body is []byte
		if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}
		return result, nil
	}
}

func tcpCall(message string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpServerAddress + ":" + tcpServerPort)
	if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
        os.Exit(1)
    }

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

	_, err = conn.Write([]byte(message))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

	println("write to server = ", message)

    conn.Close()
}

func saveSamples(fileNum int, samples Samples) {
	file, _ := files.CreateCSV("gsr" + strconv.Itoa(fileNum))
	var csvRow []string
	for i, r := range samples.GetExperimentSamples[0] {
		if (i > 1) {
			csvRow = append(csvRow, fmt.Sprintf("%f", r))
			fmt.Printf("row: %s", csvRow[0])
		}
		
	}

	csvRows := [][]string{
		{"Value"},
		csvRow,
	}

	files.WriteCSV(file, csvRows)
}
