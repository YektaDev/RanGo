package RanGo

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
