package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	d := Doller{amount: 5}
	actual := d.Times(2)

	expected := Doller{amount: 10}

	if actual != expected {
		t.Errorf("expected %v, but actual %v", expected, actual)
	}

	actual = d.Times(3)

	expected = Doller{amount: 15}

	if actual != expected {
		t.Errorf("expected %v, but actual %v", expected, actual)
	}
}
