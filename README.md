Cliarg
======

A simple way to store and modify command line executions.

NOTE: Has not been tested vigorously against untrusted input. 

Using
-----

Using `cliarg` is as simple as adding struct tags. Values are marshalled in order.

### Example Struct

```go
type ExampleCliarg struct {
	Arg         string `cliarg:"arg"` // The arg tag places the value of the field in the command line
	Option      string `cliarg:"-o,option"` // The option tag places the option before the comma in the command line, follow by the value of the field
	OptionEqual string `cliarg:"-e,optionequal"` // The optionequal tag functions similarly to the option tag, but joins the value and flag with an equal sign
	Flag        bool   `cliarg:"-f,flag"` // The flag tag either places the flag in the command line or not, based on the value of the field (bool expected)
}
```

### In Use

```go
ExampleCliarg{"/bin/arg", "option", "optionequal", true}
```

Results

    /bin/arg -o option -e=optionequal -f
