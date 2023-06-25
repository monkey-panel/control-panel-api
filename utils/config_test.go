package utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMergeObj(t *testing.T) {
	mergeValue := mergeObj(map[string]any{
		"a": "a",
		"b": "b",
		"c": []string{"test"},
		"d": "a",
		"e": map[string]any{"a": "a"},
	}, map[string]any{
		"a": "b",
		"b": "c",
		"c": "c",
		"d": []string{"test2"},
		"e": map[string]any{
			"a": "a",
			"b": "b",
			"c": []string{"test"},
			"d": "a",
		},
	})

	data, _ := json.MarshalIndent(mergeValue, "", "  ")
	// TODO add equality tests
	fmt.Printf("%v\n", string(data))

	// {
	//     "a": "b",
	//     "b": "c",
	//     "c": "c",
	//     "d": ["test2"],
	//     "e": {
	//         "a": "a",
	//         "b": "b",
	//         "c": ["test"],
	//         "d": "a"
	//     }
	// }
}
