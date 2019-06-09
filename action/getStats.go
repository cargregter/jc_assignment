package action

import (
	"encoding/json"
	"fmt"
)

// GetStats reports the statistics we can calculate from our accumulated Actions.
func GetStats() (statReport string) {
	// Summarize raw data per action in these.
	type actNumbers struct {
		count     int
		totalTime int
	}

	actSummary := make(map[string]*actNumbers)

	// Loop over the collected action detail summarizing the raw data.
	actionsSyncEle.Lock()
	for _, action := range actions {
		// If our action summary map does not yet include this action, add it.
		if _, ok := actSummary[action.Name]; !ok {
			numbers := &actNumbers{}
			actSummary[action.Name] = numbers
		}

		actSummary[action.Name].count++
		actSummary[action.Name].totalTime += action.Time
	}
	actionsSyncEle.Unlock()

	// Generate stats per action in these.
	type actionStat struct {
		Name string `json:"action"`
		Average int `json:"avg"`
	}

	// An empty slice, as opposed to a nil slice, results
	// in more intuitive output (`[]` vs. `null`) when
	// there are no recorded actions.
	actionStats := []*actionStat{}

	// Find the stats of each inputed action.
	for name, numbers := range actSummary {
		nextEntry := &actionStat{
			Name: name,
			Average: numbers.totalTime/numbers.count,
		}

		actionStats = append(actionStats, nextEntry)
	}

	// Generate our JSON string report.
	statReportBytes, err := json.MarshalIndent(actionStats, ``, `    `)
	if err != nil {
		statReport = fmt.Sprintf("error marshaling JSON string report: %s\nusing input:\n%#v",
			err.Error(), actionStats)
		return
	}

	statReport = string(statReportBytes)

	return
}