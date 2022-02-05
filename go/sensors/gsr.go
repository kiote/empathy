package sensors

import (
	"fmt"
	"net/http"
)

var (
	serverUrl = "http://localhost"
	serverPort = "22002"
	serverPath = "NeuLogAPI"
	
	startExperiment = "StartExperiment:[GSR],[1],[9],[360]"
	stopExperiment = "StopExperiment"
	getSamples = "GetExperimentSamples:[GSR],[1]"
)

func StartExperiment() {
	fmt.Printf("GSR Sensor initializing...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + startExperiment + "\n")
	httpCall(startExperiment)
}

func StopExperiment() {
	fmt.Printf("GSR Data collection stop...\n")
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + stopExperiment + "\n")
	httpCall(stopExperiment)
}

func SaveExperiment(num int) {
	fmt.Printf("GSR Data collection save...\n")
	fmt.Printf("To file %d", num)
	httpCall(getSamples)
}

func httpCall(command string) {
	fmt.Printf(serverUrl + ":" + serverPort + "/" + serverPath + "?" + command + "\n")
	resp, err := http.Get(serverUrl + ":" + serverPort + "/" + serverPath + "?" + command)
	if err != nil {
		fmt.Printf("%e \n", err)
	} else {
		fmt.Printf("%v \n", resp)
	}
}
