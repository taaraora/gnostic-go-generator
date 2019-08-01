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
				Name:      `DPaStatus`,
				Serialize: true,
			},
			result: `json:"dPaStatus,omitempty"`,
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
				Name: `DPaStatus`,
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

		tags [] string
		result       string
	}{
		{	tags: []string {`json:"createdAt,omitempty"`, `db:"updatedAt"`, `json:"verificationToken,omitempty"`},
			result:       "`"+`json:"createdAt,omitempty" db:"updatedAt" json:"verificationToken,omitempty"`+"`",
		},
		{
			tags: []string {`json:"createdAt,omitempty"`, `json:"updatedAt,omitempty"`, ``},
			result:       "`"+`json:"createdAt,omitempty" json:"updatedAt,omitempty"`+"`",
		},
		{
			tags: []string {`json:"createdAt,omitempty"`, ``, ``},
			result:       "`"+`json:"createdAt,omitempty"`+"`",
		},
		{
			tags: []string {`json:"createdAt,omitempty"`},
			result:       "`"+`json:"createdAt,omitempty"`+"`",
		},
	}
	for _, testCase := range data {
		res := structTagGenerate(testCase.tags...)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v", testCase.result, res)
		}
	}
}
