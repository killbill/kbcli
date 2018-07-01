package kbcommon

import "testing"

func TestParseLocationHeader(t *testing.T) {
	var scenarios = []struct {
		Input          string
		ExpectedResult string
	}{
		{"", ""},
		{"http://", ""},
		{"https://", ""},
		{"http://hello.com/foo/bar", "/foo/bar"},
		{"https://hello.com/foo/bar", "/foo/bar"},
		{"http://hello.com/foo/bar?hello=true", "/foo/bar?hello=true"},
		{"https://hello.com/foo/bar?hello=true", "/foo/bar?hello=true"},
	}

	for _, s := range scenarios {
		result := ParseLocationHeader(s.Input)
		if result != s.ExpectedResult {
			t.Fatalf("invalid. input: %s, expected: %s, got: %s", s.Input, result, s.ExpectedResult)
		}
	}
}
