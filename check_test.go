package ujt

import (
	"reflect"
	"testing"
)

func Test_Check(t *testing.T) {

	tests := []struct {
		name        string
		description string
		b           []byte
		want        string
	}{
		{
			name:        "1",
			description: "nil",
			b:           []byte(`null`),
			want:        "value fail",
		},
		{
			name:        "2",
			description: "empty",
			b:           []byte(`{}`),
			want:        "value fail",
		},
		{
			name:        "3",
			description: "empty",
			b:           []byte(""),
			want:        "not valid value",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			data := New(test.b, TypeDataMap)
			err := data.Check()

			if !reflect.DeepEqual(err.Error(), test.want) {
				t.Fatalf("have - %v, need - %v", err.Error(), test.want)
			}

		})
	}

}
