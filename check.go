package ujt

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TypeData string

const (
	TypeDataMap    = "map"
	TypeDataArr    = "arr"
	TypeDataCustom = "custom"
)

type Data struct {
	Value    []byte
	TypeData TypeData
}

func New(b []byte, t TypeData) *Data {
	return &Data{b, t}
}

// Check struct []byte, if not current return err
// if current unmarshal and return map[string]any
func (d *Data) Check() error {

	if !json.Valid(d.Value) {
		return errors.New("not valid value")
	}

	if len(d.Value) == 0 ||
		string(d.Value) == "{}" ||
		string(d.Value) == "[]" ||
		string(d.Value) == "null" {
		return fmt.Errorf("value fail")
	}

	return nil
}
