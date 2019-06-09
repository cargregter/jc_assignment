package action

import (
	"encoding/json"
	"fmt"
)

// Add accepts a JSON formatted string of details for an action
// and records it in memory.
func Add(actionJSON string) (err error) {
	act := &Action{}
	err = json.Unmarshal([]byte(actionJSON), act)
	if err != nil {
		err = fmt.Errorf("error unmarshaling action JSON [%s]: %s", actionJSON, err.Error())
		return
	}

	// Note: We could have made `actions` a map instead
	// and just added times as entered to a single entry
	// per distinct action. However, by keeping the actions
	// individually, we leave open (and perhaps point towards)
	// support for more functionality. This at the expense of
	// memory.
	actionsSyncEle.Lock()
	actions = append(actions, act)
	actionsSyncEle.Unlock()

	return
}
