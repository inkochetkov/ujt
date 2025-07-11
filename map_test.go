package ujt

import (
	"reflect"
	"testing"
)

func Test_UnmarshaledJSONToMap(t *testing.T) {
	tests := []struct {
		name        string
		description string
		b           []byte
		want        map[string]any
	}{
		{
			name:        "1",
			description: "data",
			b:           []byte(`{"one":"two"}`),
			want: map[string]any{
				"one": "two",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			have, err := UnmarshaledJSONToMap(test.b)
			if err != nil {
				t.Fatalf("fail %v", err)
			}

			if !reflect.DeepEqual(have, test.want) {
				t.Fatalf("have - %v, need - %v", have, test.want)
			}

		})
	}

}
