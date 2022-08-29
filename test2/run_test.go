package test2

import "testing"

func TestSayBar(t *testing.T) {
	got := SayBar()
	want := "Bar"
	if got != want {
		t.Errorf("expected: %s; got %s", want, got)
	}
}
func TestTestSelection(t *testing.T) {
	got := TestSelection()
	want := "Baz"
	if got != want {
		t.Errorf("expected: %s; got %s", want, got)
	}
}
