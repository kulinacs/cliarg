package cliarg

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
	expected string          // expected result
}{
	{TestValidCliarg{"/bin/arg", "", "", false}, "/bin/arg"},
	{TestValidCliarg{"/bin/arg", "", "", true}, "/bin/arg -f"},
	{TestValidCliarg{"/bin/arg", "", "optionequal", false}, "/bin/arg -e=optionequal"},
	{TestValidCliarg{"/bin/arg", "", "optionequal", true}, "/bin/arg -e=optionequal -f"},
	{TestValidCliarg{"/bin/arg", "option", "", false}, "/bin/arg -o option"},
	{TestValidCliarg{"/bin/arg", "option", "", true}, "/bin/arg -o option -f"},
	{TestValidCliarg{"/bin/arg", "option", "optionequal", false}, "/bin/arg -o option -e=optionequal"},
	{TestValidCliarg{"/bin/arg", "option", "optionequal", true}, "/bin/arg -o option -e=optionequal -f"},
}

func TestValidMarshal(t *testing.T) {
	for _, validCase := range validCases {
		output, _ := Marshal(validCase.cliarg)
		if output != validCase.expected {
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
