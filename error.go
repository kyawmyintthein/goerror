package errtmpl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	errorTemplate struct {
		name   string
		layout string
	}

	ErrorString struct {
		name    string
		message string
	}
)

func newErrorString(name, message string) ErrorString {
	return ErrorString{name: name, message: message}
}

func new(name, message string) error {
	return &ErrorString{name: name, message: message}
}

func (e *ErrorString) Error() string {
	return fmt.Sprintf("(#%s) %s", e.name, e.message)
}

func (e *ErrorString) IsNil() bool {
	if e.name == "" && e.message == "" {
		return true
	}
	return false
}

func (e *ErrorString) IsNotNil() bool {
	if e.IsNil() {
		return false
	}
	return true
}

func NewTemplate(name string, layout string) errorTemplate {
	return errorTemplate{name: name, layout: layout}
}

func (tmpl *errorTemplate) Error(data map[string]interface{}) error {
	return new(tmpl.name, tmpl.Parse(data))
}

func (tmpl *errorTemplate) TError(data map[string]interface{}) ErrorString {
	return newErrorString(tmpl.name, tmpl.Parse(data))
}

func (tmpl *errorTemplate) Parse(data map[string]interface{}) string {
	var layout string = tmpl.layout
	for key, val := range data {
		switch t := val.(type) {
		case string:
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), t, -1)
		case int:
			str := strconv.Itoa(t)
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), str, -1)
		default:
			s := reflect.ValueOf(t)
			layout = strings.Replace(layout, fmt.Sprintf("{{%s}}", key), fmt.Sprintf("%s", s), -1)
		}
	}
	return layout
}
