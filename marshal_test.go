package cliarg

import "reflect"
import "testing"

type TestValidCliarg struct {
	Arg         string `cliarg:"arg"`
	Option      string `cliarg:"-o"`
	Flag        bool   `cliarg:"-f"`
}

type TestInvalidCliarg struct {
	Invalid string `cliarg:"invalid"`
}

var validCases = []struct {
	cliarg   TestValidCliarg // input
	expected []string          // expected result
}{
	{TestValidCliarg{"/bin/arg", "", false}, []string{"/bin/arg"}},
	{TestValidCliarg{"/bin/arg", "", true}, []string{"/bin/arg", "-f"}},
	{TestValidCliarg{"/bin/arg", "option", false}, []string{"/bin/arg", "-o", "option"}},
	{TestValidCliarg{"/bin/arg", "option", true}, []string{"/bin/arg", "-o", "option", "-f"}},
}

func TestValidMarshal(t *testing.T) {
	for _, validCase := range validCases {
		output := Marshal(validCase.cliarg)
		if !reflect.DeepEqual(output, validCase.expected) {
			t.Errorf("Marshal was incorrect, got: %s, want: %s", output, validCase.expected)
		}
	}
}
