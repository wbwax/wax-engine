package utils

import (
	"testing"
)

func TestJoinStrWithSep(t *testing.T) {
	expected := "a=1&b=2&c=3&d=4"
	res := JoinStrWithSep("&", "a=1&b=2", "c=3", "d=4")
	if res != expected {
		t.Errorf("failed to test JoinStringWithSep, expected: %s, but got: %s", expected, res)
	}

	expected = ""
	res = JoinStrWithSep("&")
	if res != expected {
		t.Errorf("failed to test JoinStringWithSep, expected: %s, but got: %s", expected, res)
	}

	expected = "a=1"
	res = JoinStrWithSep("&", "a=1")
	if res != expected {
		t.Errorf("failed to test JoinStringWithSep, expected: %s, but got: %s", expected, res)
	}

	expected = ""
	res = JoinStrWithSep("&", "")
	if res != expected {
		t.Errorf("failed to test JoinStringWithSep, expected: %s, but got: %s", expected, res)
	}
}
