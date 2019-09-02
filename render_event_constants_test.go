package main

import (
	"log"
	"testing"
)

func Test_splitLineConstants(t *testing.T) {
	testData := []struct {
		A      string
		Result []string
	}{
		{
			A:      "- users_commands\n- accounts_commands\n- idp_commands\n",
			Result: []string{"users_commands", "accounts_commands", "idp_commands"},
		},
		{
			A:      "- users_facts\n- accounts_facts\n- idp_facts\n",
			Result: []string{"users_facts", "accounts_facts", "idp_facts"},
		},
	}
	for _, testCase := range testData {
		res := splitLineConstants(testCase.A)
		for i := 0; i < len(res); i++ {
			if res[i] != testCase.Result[i] {
				log.Printf("Expected %+q, got %+q\n", testCase.Result, res)
			}
		}
	}

}
