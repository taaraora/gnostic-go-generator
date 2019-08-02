package main

import surface "github.com/googleapis/gnostic/surface"
import "testing"

func Test_JsonTag(t *testing.T) {
	data := []struct {
		a      surface.Field
		result string
	}{
		{
			a: surface.Field{
				Name:      `DpaStatus`,
				Serialize: true,
			},
			result: `json:"dpaStatus,omitempty"`,
		},
		{
			a: surface.Field{
				Name:      ``,
				Serialize: true,
			},
			result: ``,
		},
		{
			a: surface.Field{
				Name:      ``,
				Serialize: false,
			},
			result: ``,
		},
	}
	for _, testCase := range data {
		res := jsonTag(&testCase.a)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v", testCase.result, res)
		}
	}
}
func Test_DbTeg(t *testing.T) {
	data := []struct {
		a      surface.Field
		result string
	}{
		{
			a: surface.Field{
				Name: `DpaStatus`,
			},
			result: `db:"dpa_status"`,
		},
		{
			a: surface.Field{
				Name: `dpaSTatus`,
			},
			result: `db:"dpa_status"`,
		},
		{
			a: surface.Field{
				Name: ``,
			},
			result: ``,
		},
	}
	for _, testCase := range data {
		res := dbTag(&testCase.a)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v.", testCase.result, res)
		}
	}
}

func Test_StructTagGenerate(t *testing.T) {
	data := []struct {
		testStrOne  string
		testStrTwo  string
		testStrThre string
		result      string
	}{
		{
			testStrOne:  `String one,`,
			testStrTwo:  `String two,`,
			testStrThre: `String three.`,
			result:      "` String one, String two, String three.`",
		},
		{
			testStrOne:  `String one,`,
			testStrTwo:  `String two.`,
			testStrThre: ``,
			result:      "` String one, String two. `",
		},
		{
			testStrOne:  `String one`,
			testStrTwo:  ``,
			testStrThre: ``,
			result:      "` String one  `",
		},
	}
	for _, testCase := range data {
		res := structTagGenerate(testCase.testStrOne, testCase.testStrTwo, testCase.testStrThre)
		if res != testCase.result {
			t.Fatalf("expected %v, got %v", testCase.result, res)
		}
	}
}
