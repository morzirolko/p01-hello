package greetings

import (
	"regexp"
	"testing"
)

//TestHelloName calls greetings.Hello with a name

func TestHelloName(t *testing.T) {
	name := "kekL"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("kekL")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("kekL") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with empty string

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
