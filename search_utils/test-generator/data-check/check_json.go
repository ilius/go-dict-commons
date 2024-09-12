package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := []map[string]any{}
	dataJson, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dataJson, &data)
	if err != nil {
		panic(err)
	}
	countMap := map[string]int{}
	for _, item := range data {
		b := item["b"].(string)
		score := int(item["s"].(float64))
		scoreRange := score / 10 * 10
		key := fmt.Sprintf("len: %02d, score: %02d / 20", len(b), scoreRange/10)
		countMap[key] += 1
	}
	countMapJson, err := json.MarshalIndent(countMap, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(countMapJson))
}
