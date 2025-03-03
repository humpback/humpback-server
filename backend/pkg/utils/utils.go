package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

func NewActionTimestamp() int64 {
	return time.Now().UnixMilli()
}

func PrintJson(data any) {
	value, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("%s\n", value)
}
