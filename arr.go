package ujt

import (
	"encoding/json"
	"fmt"
)

// UnmarshaledJSONToArr, in []byte, return arr
func UnmarshaledJSONToArr(b []byte) ([]any, error) {

	data := New(b, TypeDataMap)

	err := data.Check()
	if err != nil {
		return nil, err
	}

	var result []any
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, fmt.Errorf("unmarshal fail, %v", err)
	}

	return result, nil
}
