package json_merger

import (
	"reflect"
)

//var (
//	//saved struct for the interface to JSONMerger
//	jsonMerger JSONMerger
//)
//
//// The interface for the database to be used in the http handlers
//type JSONMerger interface {
//
//	// Recursively merge the "update" JSON map into the "original" JSON map
//	// @params: original - original JSON map
//	// @params: update - update JSON map
//	// @returns: merged JSON map that merged the update JSON map into the original JSON map
//	mergeJSONMaps(original, update map[string]interface{}) (interface{})
//
//	SetJSONMergerImplementation(impl JSONMerger)
//}



// Setter function for dbImplementation
//func SetJSONMergerImplementation(impl JSONMerger) {
//	jsonMerger = impl
//}

// Recursively merge the "update" JSON map into the "original" JSON map
// @params: original - original JSON map
// @params: update - update JSON map
// @returns: merged JSON map that merged the update JSON map into the original JSON map
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