package ujt

import (
	"reflect"
	"testing"
)

func Test_UnmarshaledJSONToCustom(t *testing.T) {
	type one struct {
		A string
		B string
	}

	tests := []struct {
		name        string
		description string
		b           []byte
		s           any
		want        any
	}{
		{
			name:        "1",
			description: "data",
			b:           []byte(`{"all":{"A":"one","B":"two"}}`),
			s:           one{},
			want: map[string]any{
				"all": map[string]any{
					"A": "one",
					"B": "two",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			have, err := UnmarshaledJSONToCustom(test.b, test.s)
			if err != nil {
				t.Fatalf("fail %v", err)
			}

			if !reflect.DeepEqual(have, test.want) {
				t.Fatalf("have - %v, need - %v", have, test.want)
			}

		})
	}

}
