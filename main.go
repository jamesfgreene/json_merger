package main

import (
	"reflect"
	"encoding/json"
	"fmt"
	"github.com/jamesfgreene/json_merger/json_merger"
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

	var mergedJSONMap = mergeJSONMaps(originalJSONMap, updateJSONMap)
	mergedJSONByteString, err := json.Marshal(mergedJSONMap)
	if err != nil {
		// do something here
	}

	fmt.Println("Merged JSON map/string")
	fmt.Println(mergedJSONMap)
	fmt.Println(string(mergedJSONByteString))
}



func mergeJSONMaps(original, update map[string]interface{}) (interface{}) {
	for key, value := range original {
		if _, ok := update[key]; ok == false {
			update[key] = value
		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {
			mergeJSONMaps(value.(map[string]interface{}), update[key].(map[string]interface{}))
		}
	}

	return update
}
