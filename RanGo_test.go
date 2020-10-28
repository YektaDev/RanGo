package rango

import (
	"testing"
)

func TestRnInt(t *testing.T) {
	if RnInt(0, 88) == RnInt(0, 88) {
		t.Errorf("RnInt function generates the same output each time!")
	}
	for i := 0; i < 10; i++ {
		if RnInt(0, 1) > 0 {
			t.Errorf("RnInt function generates an output greater than the greatest int allowed!")
		}
		if RnInt(0, 1) < 0 {
			t.Errorf("RnInt function generates an output smaller than the smallest int allowed!")
		}
	}
}
func TestRandomInt(t *testing.T) {
	int1, _ := RandomInt(0, 88)
	int2, _ := RandomInt(0, 88)

	if int1 == int2 {
		t.Errorf("RandomInt function generates the same output each time!")
	}
	for i := 0; i < 10; i++ {
		int3, _ := RandomInt(0, 1)

		if int3 > 0 {
			t.Errorf("RandomInt function generates an output greater than the greatest int allowed!")
		}
		if int3 < 0 {
			t.Errorf("RandomInt function generates an output smaller than the smallest int allowed!")
		}
	}
}
func TestRnString(t *testing.T) {
	if RnString(8, true, true, true, true) == RnString(8, true, true, true, true) {
		t.Errorf("RnString function generates the same output each time!")
	}

	if len(RnString(88, true, true, true, true)) != 88 {
		t.Errorf("RnString function generates an output with an unexpected length!")
	}
}
func TestRandomString(t *testing.T) {
	string1, _ := RandomString(8, true, true, true, true)
	string2, _ := RandomString(8, true, true, true, true)
	if string1 == string2 {
		t.Errorf("RandomString function generates the same output each time!")
	}

	if len(string1) != 8 {
		t.Errorf("RandomString function generates an output with an unexpected length!")
	}
}
func TestRnStringFrom(t *testing.T) {
	if RnStringFrom(8, "0123456789") == RnStringFrom(8, "0123456789") {
		t.Errorf("RnStringFrom function generates the same output each time!")
	}

	if len(RnStringFrom(88, "01")) != 88 {
		t.Errorf("RnStringFrom function generates an output with an unexpected length!")
	}
}
func TestRandomStringFrom(t *testing.T) {
	string1, _ := RandomStringFrom(8, "0123456789")
	string2, _ := RandomStringFrom(8, "0123456789")

	if string1 == string2 {
		t.Errorf("RandomStringFrom function generates the same output each time!")
	}

	if len(string1) != 8 {
		t.Errorf("RandomStringFrom function generates an output with an unexpected length!")
	}
}
