package main

import surface "github.com/googleapis/gnostic/surface"
import "testing"

func Test_JsonTag(t *testing.T) {
	data := []struct {
		field  surface.Field
		result string
	}{
		{
			field: surface.Field{
				Name:      `DpaStatus`,
				Serialize: true,
			},
			result: `json:"dpaStatus,omitempty"`,
		},
		{
			field: surface.Field{
				Name:      ``,
				Serialize: true,
			},
			result: ``,
		},
		{
			field: surface.Field{
				Name:      `DpaStatus`,
				Serialize: false,
			},
			result: ``,
		},
	}
	for _, testCase := range data {
		res := jsonTag(&testCase.field)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v", testCase.result, res)
		}
	}
}
func Test_DbTeg(t *testing.T) {
	data := []struct {
		field  surface.Field
		result string
	}{
		{
			field: surface.Field{
				Name: `DpaStatus`,
			},
			result: `db:"dpa_status"`,
		},
		{
			field: surface.Field{
				Name: `lastLoggedInIP`,
			},
			result: `db:"last_logged_in_ip"`,
		},
		{
			field: surface.Field{
				Name: ``,
			},
			result: ``,
		},
	}
	for _, testCase := range data {
		res := dbTag(&testCase.field)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v.", testCase.result, res)
		}
	}
}

func Test_StructTagGenerate(t *testing.T) {
	data := []struct {
		testTagOne   string
		testTagTwo   string
		testTagThree string
		result       string
	}{
		{
			testTagOne:   `json:"createdAt,omitempty"`,
			testTagTwo:   `json:"updatedAt,omitempty"`,
			testTagThree: `json:"verificationToken,omitempty"`,
			result:       "`"+` json:"createdAt,omitempty" json:"updatedAt,omitempty" json:"verificationToken,omitempty"`+"`",
		},
		{
			testTagOne:   `json:"createdAt,omitempty"`,
			testTagTwo:   `json:"updatedAt,omitempty"`,
			testTagThree: ``,
			result:       "`"+` json:"createdAt,omitempty" json:"updatedAt,omitempty" `+"`",
		},
		{
			testTagOne:   `json:"createdAt,omitempty"`,
			testTagTwo:   ``,
			testTagThree: ``,
			result:       "`"+` json:"createdAt,omitempty"  `+"`",
		},
	}
	for _, testCase := range data {
		res := structTagGenerate(testCase.testTagOne, testCase.testTagTwo, testCase.testTagThree)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v", testCase.result, res)
		}
	}
}
