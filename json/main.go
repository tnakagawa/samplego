// json project main.go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("BEGIN")
	jsonMap := make(map[string]interface{})
	jsonStr := "{\"aaa\":\"bbb\",\"ccc\":123}"
	json.Unmarshal([]byte(jsonStr), &jsonMap)
	jsonBs, _ := json.MarshalIndent(jsonMap, "", "\t")
	fmt.Println(string(jsonBs))
	fmt.Println("END")
}
