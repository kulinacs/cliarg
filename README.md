Cliarg
======

[![Build Status](https://travis-ci.org/kulinacs/cliarg.svg?branch=master)](https://travis-ci.org/kulinacs/cliarg)
[![Coverage Status](https://coveralls.io/repos/github/kulinacs/cliarg/badge.svg?branch=master)](https://coveralls.io/github/kulinacs/cliarg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kulinacs/cliarg)](https://goreportcard.com/report/github.com/kulinacs/cliarg)
[![GoDoc](https://godoc.org/github.com/kulinacs/cliarg?status.svg)](https://godoc.org/github.com/kulinacs/cliarg)

A simple way to store and modify command line executions.

Using
-----

Using `cliarg` is as simple as adding struct tags. Values are marshalled in order.

### Example Struct

```go
type ExampleCliarg struct {
	Arg         string `cliarg:"arg"` // The arg tag places the value of the field in the command line
	Option      string `cliarg:"-o"`  // An option tag places the tag on the command line, followed by the value of the field
	Flag        bool   `cliarg:"-f"`  // A flag tag either places the flag in the command line or not, based on the value of field, occurs if boolean
}
```

### In Use

```go
ExampleCliarg{"/bin/arg", "option", true}
```

Results

    /bin/arg -o option -f
