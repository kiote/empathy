package sensors

import (
	"fmt"
	"net/http"
)

var (
	serverUrl = "http://localhost"
	serverPort = "22002"
	serverPath = "NeuLogAPI"
	
	startExperiment = "StartExperiment:[GSR],[1],[9],[600]"
	stopExperiment = "StopExperiment"
	getSamples = "GetExperimentSamples:[GSR],[1]"
)

func StartExperiment() {
	fmt.Printf("GSR Sensor initializing...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + startExperiment + "\n")
	resp, err := http.Get(serverUrl + ":" + serverPort + "/" + serverPath + "?" + startExperiment)
	if err != nil {
		fmt.Printf("%e \n", err)
	} else {
		fmt.Printf("%v \n", resp)
	}
}

func StopExperiment() {
	fmt.Printf("GSR Data collection stop...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + stopExperiment + "\n")
	resp, err := http.Get(serverUrl + ":" + serverPort + "/" + serverPath + "?" + stopExperiment)
	if err != nil {
		fmt.Printf("%e \n", err)
	} else {
		fmt.Printf("%v \n", resp)
	}
}

func SaveExperiment(num int) {
	fmt.Printf("GSR Data collection save...\n")
	fmt.Printf("To file %d", num)
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + stopExperiment + "\n")
	resp, err := http.Get(serverUrl + ":" + serverPort + "/" + serverPath + "?" + getSamples)
	if err != nil {
		fmt.Printf("%e \n", err)
	} else {
		fmt.Printf("%v \n", resp)
	}
}
