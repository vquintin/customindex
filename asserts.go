package customindex

import "testing"

func assertEquals(t *testing.T, message string, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("%v. Expected: %v. Got: %v.", message, expected, actual)
	}
}

func assertNil(t *testing.T, message string, v interface{}) {
	if v != nil {
		t.Errorf("%v. Got: %v", message, v)
	}
}

func assertNoError(t *testing.T, err error) {
	assertNil(t, "An error occured", err)
}

func assertFalse(t *testing.T, message string, b bool) {
	assertTrue(t, message, !b)
}

func assertTrue(t *testing.T, message string, b bool) {
	if !b {
		t.Errorf("%v", message)
	}
}
