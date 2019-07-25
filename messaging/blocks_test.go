package messaging

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBlocksEncoding(t *testing.T) {
	cases := []struct {
		desc string
		in   interface{}
		out  map[string]interface{}
	}{
		{
			desc: "Section",
			in: Section{
				Text: "This is the required text for the block",
			},
			out: map[string]interface{}{
				"type": "section",
				"text": "This is the required text for the block",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			d, err := json.Marshal(c.in)

			if err != nil {
				t.Errorf("unexpected error %q", err)
			}

			var actual map[string]interface{}

			json.Unmarshal(d, &actual)

			if diff := cmp.Diff(c.out, actual); diff != "" {
				t.Error("unexpected diff\n", diff)
			}
		})
	}
}
