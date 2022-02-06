package settings

import (
    "fmt"
	"time"
)

var (
	CurrentTimestamp = time.Now().Unix()
	DataDir = "../" +  fmt.Sprintf("%d", CurrentTimestamp)
)