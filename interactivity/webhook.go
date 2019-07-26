package interactivity

import (
	"encoding/json"
	"net/http"

	"github.com/geckoboard/slackutil-go/messaging"
)

type Request struct {
	// this helps identify which type of interactive component sent the
	// payload; An interactive element in a block will have a type of
	// `block_actions`, whereas an interactive element in an attachment will
	// have a type of `interactive_message`.
	Type        string               `json:"type"`
	TriggerID   string               `json:"trigger_id"`
	ResponseURL string               `json:"response_url"`
	Actions     []StaticSelectAction `json:"actions"`
}

// Action contains data from the specific interactive component that was used
type Action struct {
	// this will identify the block within a message that contained the
	// interactive component that was used
	BlockID string `json:"block_id"`
	// this identifies the interactive component itself. Some blocks can
	// contain multiple interactive components, so the block_id alone may
	// not be specific enough to identify the source component
	ActionID string `json:"action_id"`
	// set by your app when you composed the message, this is the value
	// that was specified in the interactive component when an interaction
	// happened. For example, a select menu will have multiple possible
	// values depending on what the user picks from the menu, and value
	// will identify the chosen option
	Value string `json:"value"`
}

type StaticSelectAction struct {
	Action
	SelectedOption messaging.MenuOption `json:"selected_option"`
}

type MessageResponder struct {
	r Request
}

func Handler(with func(Request, MessageResponder)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sr Request

		json.Unmarshal([]byte(r.FormValue("payload")), &sr)

		with(sr, MessageResponder{sr})
	}
}
