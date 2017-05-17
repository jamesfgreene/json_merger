package merger

import (
	"reflect"
)

type Merger interface {
	MergeJSONMaps(original, update map[string]interface{}) (interface{})
}

type Merge struct {
}

// Recursively merge the "update" JSON map into the "original" JSON map
// @params: original - original JSON map
// @params: update - update JSON map
// @returns: merged JSON map that merged the update JSON map into the original JSON map
func (m Merge) MergeJSONMaps(original, update map[string]interface{}) (interface{}) {
	for key, value := range original {
		if _, ok := update[key]; ok == false {
			update[key] = value
		} else if reflect.TypeOf(value).String() == "map[string]interface {}" {
			m.MergeJSONMaps(value.(map[string]interface{}), update[key].(map[string]interface{}))
		}
	}

	return update
}

func NewMerger() Merger {
	return &Merge{}
}