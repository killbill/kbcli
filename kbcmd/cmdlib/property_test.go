package cmdlib

import (
	"testing"

	"github.com/go-openapi/strfmt"

	"github.com/google/go-cmp/cmp"
)

func TestParseProperties_Valid(t *testing.T) {
	result, err := ParseProperties([]string{"one=1", "two=2"})
	if err != nil {
		t.Fatal(err)
	}
	exp := map[string]string{
		"one": "1",
		"two": "2",
	}
	if diff := cmp.Diff(exp, result); diff != "" {
		t.Fatal(diff)
	}
}

func TestParseProperties_Quoted(t *testing.T) {
	result, err := ParseProperties([]string{
		`one=hello one`,
		`foo=foo bar`,
	})
	if err != nil {
		t.Fatal(err)
	}
	exp := map[string]string{
		"one": "hello one",
		"foo": "foo bar",
	}
	if diff := cmp.Diff(exp, result); diff != "" {
		t.Fatal(diff)
	}
}

type testObj struct {
	AccountID          string
	ParentID           string
	CompanyName        *string
	XKillbillAPIKey    string
	XKillbillAPISecret string
	IsDefault          bool
	IsDefaultPtr       *bool
	UniqueID           strfmt.UUID
	UniqueIDPtr        *strfmt.UUID
}

func TestLoadProperties(t *testing.T) {
	uuid1 := strfmt.UUID("12345678-1234-1234-1234-123456789012")
	uuid2 := strfmt.UUID("23456789-2345-2345-2345-234567890123")
	kv := map[string]string{
		"AccountID":    "123",
		"ParentId":     "333",
		"companyname":  "google",
		"isdefault":    "true",
		"isdefaultPtr": "false",
		"uniqueid":     string(uuid1),
		"uniqueIdPtr":  string(uuid2),
	}
	obj := testObj{}
	err := LoadProperties(kv, &obj)
	if err != nil {
		t.Fatal(err)
	}

	companyName := "google"
	isDefault := false
	exp := testObj{
		AccountID:    "123",
		ParentID:     "333",
		CompanyName:  &companyName,
		IsDefault:    true,
		IsDefaultPtr: &isDefault,
		UniqueID:     strfmt.UUID(uuid1),
		UniqueIDPtr:  &uuid2,
	}
	if diff := cmp.Diff(exp, obj); diff != "" {
		t.Fatal(diff)
	}
}

func TestListProperties(t *testing.T) {
	result := ListProperties(&testObj{})
	exp := []string{"AccountID", "ParentID", "CompanyName", "IsDefault", "IsDefaultPtr", "UniqueID", "UniqueIDPtr"}
	if diff := cmp.Diff(exp, result); diff != "" {
		t.Fatal(diff)
	}
}

func TestGenerateUsageString_Simple(t *testing.T) {
	result := GenerateUsageString(&testObj{}, []string{"+accountid", "-companyName"})
	exp := []string{"AccountID=string", "[ParentID=string]", "[IsDefault=bool]", "[IsDefaultPtr=*bool]", "[UniqueID=strfmt.UUID]", "[UniqueIDPtr=*strfmt.UUID]"}
	if diff := cmp.Diff(exp, result); diff != "" {
		t.Fatal(diff)
	}
}

func TestGenerateUsageString_Full(t *testing.T) {
	result := GenerateUsageString(&testObj{}, []string{})
	exp := []string{"[AccountID=string]", "[ParentID=string]", "[CompanyName=*string]", "[IsDefault=bool]", "[IsDefaultPtr=*bool]", "[UniqueID=strfmt.UUID]", "[UniqueIDPtr=*strfmt.UUID]"}
	if diff := cmp.Diff(exp, result); diff != "" {
		t.Fatal(diff)
	}
}
