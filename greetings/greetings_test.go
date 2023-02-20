package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Prashant"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Prashant")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Prashant") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
