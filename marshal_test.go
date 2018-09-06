package cliarg

import "reflect"
import "testing"

type TestValidCliarg struct {
	Arg         string `cliarg:"arg"`
	Option      string `cliarg:"-o,option"`
	OptionEqual string `cliarg:"-e,optionequal"`
	Flag        bool   `cliarg:"-f,flag"`
}

type TestInvalidCliarg struct {
	Invalid string `cliarg:"invalid"`
}

var validCases = []struct {
	cliarg   TestValidCliarg // input
	expected []string          // expected result
}{
	{TestValidCliarg{"/bin/arg", "", "", false}, []string{"/bin/arg"}},
	{TestValidCliarg{"/bin/arg", "", "", true}, []string{"/bin/arg", "-f"}},
	{TestValidCliarg{"/bin/arg", "", "optionequal", false}, []string{"/bin/arg", "-e=optionequal"}},
	{TestValidCliarg{"/bin/arg", "", "optionequal", true}, []string{"/bin/arg", "-e=optionequal", "-f"}},
	{TestValidCliarg{"/bin/arg", "option", "", false}, []string{"/bin/arg", "-o", "option"}},
	{TestValidCliarg{"/bin/arg", "option", "", true}, []string{"/bin/arg", "-o", "option", "-f"}},
	{TestValidCliarg{"/bin/arg", "option", "optionequal", false}, []string{"/bin/arg", "-o", "option", "-e=optionequal"}},
	{TestValidCliarg{"/bin/arg", "option", "optionequal", true}, []string{"/bin/arg", "-o", "option", "-e=optionequal", "-f"}},
}

func TestValidMarshal(t *testing.T) {
	for _, validCase := range validCases {
		output, _ := Marshal(validCase.cliarg)
		if !reflect.DeepEqual(output, validCase.expected) {
			t.Errorf("Marshal was incorrect, got: %s, want: %s", output, validCase.expected)
		}
	}
}

func TestInvalidMarshal(t *testing.T) {
	testValue := TestInvalidCliarg{""}
	_, err := Marshal(testValue)
	if err == nil {
		t.Errorf("Function was expected to throw an error")
	}
}
