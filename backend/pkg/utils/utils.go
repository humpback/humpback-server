package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data any) {
	value, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("%s\n", value)
}
