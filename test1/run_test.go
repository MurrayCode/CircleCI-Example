package test1

import "testing"

func TestSayFoo(t *testing.T) {
	got := SayFoo()
	want := "Foo"
	if got != want {
		t.Errorf("expected: %s; got %s", want, got)
	}
}
