package main

import (
	"encoding/json"
	"fmt"
	"github.com/jamesfgreene/json_merger/merger"
)

func main() {
	var originalJSONMap map[string]interface{}
	json.Unmarshal([]byte(`{"foo":{"bop":{"boi":"goi"}}, "baz":"boo", "haz":[1,2,3]}`), &originalJSONMap)
	originalJSONByteString, err := json.Marshal(originalJSONMap)
	if err != nil {
		// do some logging here
	}

	var updateJSONMap map[string]interface{}
	json.Unmarshal([]byte(`{"foo":{"bop":{"toi":"qoi", "boi":"boom"}}, "gaz":"goo", "taz":[4,5,6], "haz":[8,9,10]}`), &updateJSONMap)
	updateJSONByteString, err := json.Marshal(updateJSONMap)
	if err != nil {
		// do some logging here
	}

	fmt.Println("First JSON map/string")
	fmt.Println(originalJSONMap)
	fmt.Println(string(originalJSONByteString)+"\n")

	fmt.Println("Second JSON map/string")
	fmt.Println(updateJSONMap)
	fmt.Println(string(updateJSONByteString)+"\n")

	jsonMerger := merger.NewMerger()


	var mergedJSONMap = jsonMerger.MergeJSONMaps(originalJSONMap, updateJSONMap)
	mergedJSONByteString, err := json.Marshal(mergedJSONMap)
	if err != nil {
		// do something here
	}

	fmt.Println("Merged JSON map/string")
	fmt.Println(mergedJSONMap)
	fmt.Println(string(mergedJSONByteString))
}