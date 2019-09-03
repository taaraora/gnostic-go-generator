package main

import (
	"errors"
	"testing"
)

func Test_splitLineConstants(t *testing.T) {
	testData := []struct {
		A      string
		name   string
		Result []string
		err    error
	}{
		{
			A:      "- users_commands\n- accounts_commands\n- idp_commands\n",
			name:   "",
			Result: []string{"users_commands", "accounts_commands", "idp_commands"},
			err:    error(nil),
		},
		{
			A:      "- users_facts\n- accounts_facts\n- idp_facts\n",
			name:   "",
			Result: []string{"users_facts", "accounts_facts", "idp_facts"},
			err:    error(nil),
		},
		{
			A:      "properties:\n  code:\n    type: integer\n    format: int64\n    message:\n    type: string\n",
			name:   "ls",
			Result: []string{"properties:", "  code:", "    type: integer", "    format: int64", "    message:", "    type: string"},
			err: errors.New(`Specification Extensions : ` + `ls` + `
the specification can be expanded, for the generation of constants, only a slice of strings.

example:
         x-constants:
           - constant_1
           - constant_2  
           ...`),
		},
	}
	for _, testCase := range testData {
		res, err := splitLineConstants(testCase.A, testCase.name)
		if err != nil {

		}
		for i := 0; i < len(res); i++ {
			if res[i] != testCase.Result[i] {
				t.Errorf("Expected %+q, got %+q\n", testCase.Result, res)
			}
			if err != nil && err.Error() != testCase.err.Error() {
				t.Errorf("Expected %+v, got %+v", testCase.err.Error(), err.Error())
			}
		}
	}

}
