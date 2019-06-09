package action

import(
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestAdd provides unit testing for function Add.
func TestAdd(t *testing.T) {
	Convey("Unmarshaling fails.\n", t, func() {
		badJSON := "badJSON"

		// Function under test.
		err := Add(badJSON)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual,
			`error unmarshaling action JSON [badJSON]: invalid character 'b' looking for beginning of value`)
		So(len(actions), ShouldEqual, 0)
	})

	Convey("Adding one action succeeds.\n", t, func() {
		// Function under test.
		err := Add(action1)

		So(err, ShouldBeNil)
		So(len(actions), ShouldEqual, 1)
	})

	Convey("Adding multiple actions in parallel succeeds.\n", t, func() {
		// Kick off one client.
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

		// Kick off another client.
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

		// Wait for both clients to finish.
		<- c1
		<- c2

		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(len(actions), ShouldEqual, 21) // 1 added in the previous test case.
	})

	// Clean up from these tests.
	actions = nil
}
