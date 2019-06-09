package action

import "sync"

// Action is used to capture the details for a single action.
type Action struct {
	Name string `json:"action"`
	Time int `json:"time"`
}

// Actions is a collection of Action
type Actions []*Action

// actions captures our input Actions.
var actions Actions

// actionsSyncEle is used to synchronize access to actions.
var actionsSyncEle = &sync.Mutex{}
