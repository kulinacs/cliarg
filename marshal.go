package cliarg

import (
	"reflect"
)

// Marshal returns the command line encoding of v.
func Marshal(v interface{}) []string {
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	var output []string
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("cliarg")
		if tag == "arg" {
			output = append(output, reflect.Indirect(val).FieldByName(field.Name).String())
		} else {
			if field.Type.Kind() == reflect.Bool {
				if reflect.Indirect(val).FieldByName(field.Name).Bool() {
					output = append(output, tag)
				}
			} else {
				option := reflect.Indirect(val).FieldByName(field.Name).String()
				if len(option) > 0 {
					output = append(output, tag)
					output = append(output, option)
				}
			}
		}
	}
	return output
}
