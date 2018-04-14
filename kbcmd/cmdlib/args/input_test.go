package args

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProperties_Get(t *testing.T) {
	var properties = Properties([]Property{
		{Name: "P1"},
		{Name: "P2"},
	})
	if properties.Get("p1").Name != "P1" {
		t.Fatalf("can't get property p1")
	}
	if properties.Get("P2").Name != "P2" {
		t.Fatalf("can't get property p2")
	}
}

func TestProperties_Remove(t *testing.T) {
	var properties = Properties([]Property{
		{Name: "P1"},
		{Name: "P2"},
		{Name: "P3"},
	})
	properties.Remove("P2")
	exp := Properties([]Property{
		{Name: "P1"},
		{Name: "P3"},
	})
	if diff := cmp.Diff(exp, properties); diff != "" {
		t.Fatal(diff)
	}

	properties.Remove("P1")
	exp = Properties([]Property{
		{Name: "P3"},
	})
	if diff := cmp.Diff(exp, properties); diff != "" {
		t.Fatal(diff)
	}

	properties.Remove("P3")
	exp = Properties([]Property{})
	if diff := cmp.Diff(exp, properties); diff != "" {
		t.Fatal(diff)
	}
}

func TestInputArg_Split(t *testing.T) {
	testConfigs := []struct {
		Input string
		Key   string
		Value string
		Error string
	}{
		{"foo=bar", "foo", "bar", ""},
		{"foo=foo bar", "foo", "foo bar", ""},
		{"foo=", "foo", "", ""},
		{"foo", "", "", "Invalid input. Expecting PROPERTY=VALUE"},
	}

	for _, tc := range testConfigs {
		k, v, err := (InputArg(tc.Input)).Split()
		var errStr string
		if err != nil {
			errStr = err.Error()
		}
		if k != tc.Key || v != tc.Value || errStr != tc.Error {
			t.Fatalf("invalid value. expecting %s,%s,%s, got %s,%s,%s", tc.Key, tc.Value, tc.Error, k, v, errStr)
		}
	}
}
