package assert

import "testing"

func AssertEquals(t *testing.T, message string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("%v. Expected: %v. Got: %v.", message, expected, actual)
	}
}

func AssertNil(t *testing.T, message string, v interface{}) {
	if v != nil {
		t.Errorf("%v. Got: %v", message, v)
	}
}

func AssertNoError(t *testing.T, err error) {
	AssertNil(t, "An error occured", err)
}

func AssertFalse(t *testing.T, message string, b bool) {
	AssertTrue(t, message, !b)
}

func AssertTrue(t *testing.T, message string, b bool) {
	if !b {
		t.Errorf("%v", message)
	}
}
