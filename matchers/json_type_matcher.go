package matchers

import (
	"fmt"
	"reflect"

	"github.com/totherme/nosj"
)

// Deprecated: use gnosj.JSONTypeMatcher instead
type JSONTypeMatcher struct {
	typ string
}

// Deprecated: use gnosj.BeAnObject instead
func BeAnObject() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONOb,
	}
}

// Deprecated: use gnosj.BeAString instead
func BeAString() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONString,
	}
}

// Deprecated: use gnosj.BeAList instead
func BeAList() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONList,
	}
}

// Deprecated: use gnosj.BeANum instead
func BeANum() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONNum,
	}
}

// Deprecated: use gnosj.BeABool instead
func BeABool() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONBool,
	}
}

// Deprecated: use gnosj.BeANull instead
func BeANull() JSONTypeMatcher {
	return JSONTypeMatcher{
		typ: nosj.JSONNull,
	}
}

// Deprecated: use gnosj.JSONTypeMatcher instead
func (m JSONTypeMatcher) Match(actual interface{}) (success bool, err error) {
	switch json := actual.(type) {
	default:
		return false, fmt.Errorf("actual is not a JSON -- actually of type %s", reflect.TypeOf(actual))
	case nosj.JSON:
		return json.IsOfType(m.typ), nil
	}
}

// Deprecated: use gnosj.JSONTypeMatcher instead
func (m JSONTypeMatcher) FailureMessage(actual interface{}) (message string) {
	if reflect.TypeOf(actual) != reflect.TypeOf(nosj.JSON{}) {
		return fmt.Sprintf("expected a JSON object -- got a %s", reflect.TypeOf(actual))
	}

	json := actual.(nosj.JSON)
	for _, t := range []string{nosj.JSONBool,
		nosj.JSONString,
		nosj.JSONNum,
		nosj.JSONList,
		nosj.JSONNull,
		nosj.JSONOb} {
		if json.IsOfType(t) {
			return fmt.Sprintf("expected a JSON %s -- got a JSON %s", m.typ, t)
		}
	}

	return fmt.Sprintf("expected a JSON %s -- got some other crazy kind of JSON", m.typ)
}

// Deprecated: use gnosj.JSONTypeMatcher instead
func (m JSONTypeMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("got a JSON %s, but expected not to", m.typ)
}
