package cliarg

import "errors"
import "fmt"
import "reflect"
import "strings"

// Marshal returns the command line encoding of v.
func Marshal(v interface{}) (string, error) {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	var output []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := strings.Split(field.Tag.Get("cliarg"), ",")
		switch tag[len(tag)-1] {
		case "arg":
			output = append(output, reflect.Indirect(val).FieldByName(field.Name).String())
		case "option":
			option := reflect.Indirect(val).FieldByName(field.Name).String()
			if len(option) > 0 {
				output = append(output, fmt.Sprintf("%v %v", tag[0], option))
			}
		case "optionequal":
			optionEqual := reflect.Indirect(val).FieldByName(field.Name).String()
			if len(optionEqual) > 0 {
				output = append(output, fmt.Sprintf("%v=%v", tag[0], optionEqual))
			}
		case "flag":
			if reflect.Indirect(val).FieldByName(field.Name).Bool() {
				output = append(output, fmt.Sprintf("%v", tag[0]))
			}
		default:
			return "", errors.New("unknown struct tag type")
		}
	}
	return strings.Join(output, " "), nil
}
