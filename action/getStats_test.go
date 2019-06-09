package action

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestGetStatistics provides unit testing for function GetStats.
func TestGetStatistics(t *testing.T) {
	Convey("No actions.\n", t, func() {
		// Function under test.
		report := GetStats()

		So(report, ShouldEqual, `[]`)
	})

	Convey("One action.\n", t, func() {
		act := &Action{}
		err := json.Unmarshal([]byte(action1), act)
		if err != nil {
			t.Fatalf("error unmarshaling action JSON [%s]: %s", action1, err.Error())
		}

		actions = append(actions, act)

		// Function under test.
		report := GetStats()

		So(report, ShouldEqual, `[
    {
        "action": "jump",
        "avg": 100
    }
]`)
	})

	Convey("Several actions, some of different names.\n", t, func() {
		act := &Action{}
		err := json.Unmarshal([]byte(action2), act)
		if err != nil {
			t.Fatalf("error unmarshaling action JSON [%s]: %s", action1, err.Error())
		}

		actions = append(actions, act)

		act = &Action{}
		err = json.Unmarshal([]byte(action3), act)
		if err != nil {
			t.Fatalf("error unmarshaling action JSON [%s]: %s", action1, err.Error())
		}

		actions = append(actions, act)

		if len(actions) != 3 {
			t.Fatalf("unexpected number of actions [%d]: %#v", len(actions), actions)
		}

		// Function under test.
		report := GetStats()

		// No order to JSON lists.
		// ShouldEqual() returns "" when successful.
		if ShouldEqual(report, `[
    {
        "action": "jump",
        "avg": 150
    },
    {
        "action": "run",
        "avg": 75
    }
]`) != "" {
			if ShouldEqual(report, `[
    {
        "action": "run",
        "avg": 75
    },
    {
        "action": "jump",
        "avg": 150
    }
]`) != "" {
			}
			So(report, ShouldEqual, `[
    {
        "action": "jump",
        "avg": 150
    },
    {
        "action": "run",
        "avg": 75
    }
]`)
		}
	})

//	Convey("Marshaling fails.\n", t, func() {
//		We would have to replace a direct call to json.MarshalIndent with some
//      interface, generate a mock for the interface and inject it in this case,
//      having it return an error.
//	})

	// Clean up from these tests.
	actions = nil

}
