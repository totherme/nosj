package matchers

import (
	"fmt"
	"github.com/totherme/nosj"
)

// Deprecated: use gnosj.HaveJSONPointerMatcher instead
type HaveJSONPointerMatcher struct {
	p string
}

// Deprecated: use gnosj.HaveJSONPointer instead
func HaveJSONPointer(p string) HaveJSONPointerMatcher {
	return HaveJSONPointerMatcher{p: p}
}

// Deprecated: use gnosj.HaveJSONPointerMatcher instead
func (m HaveJSONPointerMatcher) Match(actual interface{}) (bool, error) {

	switch t := actual.(type) {
	default:
		return false, fmt.Errorf("not a JSON object. Have you done nosj.ParseJSON(...)?")
	case nosj.JSON:
		return t.HasPointer(m.p)
	}
}

// Deprecated: use gnosj.HaveJSONPointerMatcher instead
func (m HaveJSONPointerMatcher) FailureMessage(actual interface{}) (message string) {
	actualString := fmt.Sprintf("%+v", actual)
	return fmt.Sprintf("expected '%s' to be a nosj.JSON object with pointer '%s'",
		truncateString(actualString),
		m.p)
}

// Deprecated: use gnosj.HaveJSONPointerMatcher instead
func (m HaveJSONPointerMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	actualString := fmt.Sprintf("%+v", actual)
	return fmt.Sprintf("expected '%s' not to contain the pointer '%s'",
		truncateString(actualString),
		m.p)
}
