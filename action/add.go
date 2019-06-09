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

	actionsSyncEle.Lock()
	actions = append(actions, act)
	actionsSyncEle.Unlock()

	return
}
