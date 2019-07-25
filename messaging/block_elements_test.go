package messaging

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBlockElementsEncoding(t *testing.T) {
	cases := []struct {
		desc string
		in   interface{}
		out  map[string]interface{}
	}{
		{
			desc: "Static select",
			in: StaticSelect{
				Placeholder: PlainText("This shows up if selects are not supported"),
				Options: []MenuOption{
					{
						Text:  PlainText("This is what the user sees"),
						Value: "This is sent to your app",
					},
				},
			},
			out: map[string]interface{}{
				"type":        "static_select",
				"placeholder": "This shows up if selects are not supported",
				"action_id":   "",
				"options": []interface{}{
					map[string]interface{}{
						"text":  "This is what the user sees",
						"value": "This is sent to your app",
					},
				},
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
