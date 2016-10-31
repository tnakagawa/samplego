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

	prop()

	fmt.Println("END")
}

var propMap = make(map[string]interface{})

func prop() {
	propMap["aaa"] = 1
	propMap["bbb"] = "ccc"
	mapItem := make(map[string]string)
	mapItem["123"] = "456"
	mapItem["456"] = "789"
	propMap["ccc"] = mapItem

	jsonBs, _ := json.MarshalIndent(propMap, "", "    ")
	fmt.Println(string(jsonBs))

	vs, ok := propMap["aaa"].(int)
	fmt.Println(vs, ok)

	v, ok := propMap["bbb"].(string)
	fmt.Println(v, ok)

}

func set(key string, value interface{}) {

}
