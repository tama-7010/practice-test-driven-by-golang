package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	got := Doller{amount: 5}
	got.Times(2)

	want := Doller{amount: 10}

	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}
}
