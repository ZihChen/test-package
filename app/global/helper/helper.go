package helper

import (
	"encoding/json"
	"fmt"
)

func PrintLog(param interface{}) {
	byteData, _ := json.Marshal(param)
	fmt.Println(string(byteData))
}