package action

import(
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// TestAddAndGet provides integration testing for functions Add and GetStats.
// This test could be made a bit more comprehensive - e.g., testing intermittent reports.
// But due to the parallelism allowed by Add() and GetStats(), this improvement,
// without some care, could yield "flakiness".
func TestAddAndGet(t *testing.T) {
	Convey("Adding multiple actions in parallel to getting statistics succeeds.\n", t, func() {
		// Kick off one add client.
		c1 := make(chan error)
		var err1 error
		go func() {
			// Notify the caller when we're done, no matter how we end up.
			defer func() {c1 <- err1}()

			for i := 0; i < 10; i++ {
				// Function under test.
				err1 = Add(action2)
				if err1 != nil {
					break
				}
			}
		}()

		// Kick off another add client.
		c2 := make(chan error)
		var err2 error
		go func() {
			// Notify the caller when we're done, no matter how we end up.
			defer func() {c2 <- err2}()

			for i := 0; i < 10; i++ {
				// Function under test.
				err2 = Add(action3)
				if err2 != nil {
					break
				}
			}
		}()

		// Kick off a get client.
		c3 := make(chan string)
		var report string
		go func() {
			// Notify the caller when we're done, no matter how we end up.
			defer func() {c3 <- report}()

			for i := 0; i < 5; i++ {
				// Function under test.
				report = GetStats()
				time.Sleep(time.Microsecond)
			}
		}()

		// Wait for clients to finish.
		<- c1
		<- c2
		<- c3

		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(len(actions), ShouldEqual, 20)

		// One last request of the report.
		report = GetStats()

		// No order to JSON lists.
		// ShouldEqual() returns "" when successful.
		if ShouldEqual(report, `[
    {
        "action": "jump",
        "avg": 200
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
        "avg": 200
    }
]`) != "" {
				So(report, ShouldEqual, `[
    {
        "action": "jump",
        "avg": 200
    },
    {
        "action": "run",
        "avg": 75
    }
]`)
			}
		}
})

	// Clean up from these tests.
	actions = nil
}
