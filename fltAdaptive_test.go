package transforms

import (
	"strconv"
	"testing"
)

// FltSlice tests with negative integers
//---------------------------------------------------------------------------------------------------- <-100

func ExampleTestAdaptiveZoneFlt() {
	AdaptiveZoneFlt(0.50, 1.16724, 80.0, 20.0)
}

func TestAdaptiveZoneFlt(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log("Testing: positive [n] values for AdaptiveZone")
		y := float64(i) / 100.0
		n := AdaptiveZoneFlt(y, 1, 80.0, 100.0)
		if n != -1 {
			t.Error("AdaptiveZoneFlt() not expected value of -1: " + strconv.FormatFloat(n, 'f', 2, 64))
		}
	}
}
