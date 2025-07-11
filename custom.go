package ujt

import (
	"encoding/json"
	"fmt"
)

// UnmarshaledJSONToCustom, in []byte, return custom struct
func UnmarshaledJSONToCustom(b []byte, customStruct any) (any, error) {

	data := New(b, TypeDataMap)

	err := data.Check()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &customStruct)
	if err != nil {
		return nil, fmt.Errorf("unmarshal fail, %v", err)
	}

	return customStruct, nil
}
