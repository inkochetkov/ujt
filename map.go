package ujt

import (
	"encoding/json"
	"fmt"
)

// UnmarshaledJSONToMap, in []byte, return map
func UnmarshaledJSONToMap(b []byte) (map[string]any, error) {

	data := New(b, TypeDataMap)

	err := data.Check()
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal fail, %v", err)
	}

	return result, nil
}
