package hello

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Alex"
	want := regexp.MustCompile("Hello " + name)

	got := Hello(name)
	if !want.MatchString(got) {
		t.Fatalf("Hello(%s) = %s != %s", name, got, want)
	}
}
