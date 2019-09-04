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
			input:  "properties:\n  code:\n    type: integer\n    format: int64\n    message:\n    type: string\n",
			result: nil,
			err:    errors.New("can't parse yaml list"),
		},
	}
	for _, testCase := range testData {
		res, err := splitLineConstants(testCase.input)
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("expected :%+q:, got :%+q:", err, testCase.err)
		}
		if !reflect.DeepEqual(res, testCase.result) {
			t.Errorf("expected %+v, got %+v", res, testCase.result)
		}
	}
}

func Test_generateStreamTypeConstants(t *testing.T) {
	testData := []struct {
		input  string
		result []string
		err    error
	}{
		{
			input:  "- users_commands\n- accounts_commands\n- idp_commands\n",
			result: []string{"const (", "UsersCommandsStreamType = \"users_commands\"", "AccountsCommandsStreamType = \"accounts_commands\"", "IdpCommandsStreamType = \"idp_commands\"", ")"},
			err:    nil,
		},
		{
			input:  "- users_facts\n- accounts_facts\n- idp_facts\n",
			result: []string{"const (", "UsersFactsStreamType = \"users_facts\"", "AccountsFactsStreamType = \"accounts_facts\"", "IdpFactsStreamType = \"idp_facts\"", ")"},
			err:    nil,
		},
		{
			input:  "properties:\n  code:\n    type: integer\n    format: int64\n    message:\n    type: string\n",
			result: nil,
			err:    errors.New("can't parse yaml list"),
		},
	}
	for _, testCase := range testData {
		res, err := generateStreamTypeConstants(testCase.input)
		if err != nil && err.Error() != testCase.err.Error() {
			t.Errorf("expected :%+q:, got :%+q:", err, testCase.err)
		}
		if !reflect.DeepEqual(res, testCase.result) {
			t.Errorf("expected %+q, got %+q", res, testCase.result)
		}
	}
}
