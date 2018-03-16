package transforms

import (
	"strconv"
	"testing"
)

// FltSlice tests with negative integers
//-------------------------------------------------------------------------------------------------- <-100

func ExampleAdaptiveZoneFlt() {
	AdaptiveZoneFlt(0.99, 1.16724, 80.0, 20.0)
}

func TestAdaptiveZoneFlt(t *testing.T) {
	t.Log("Testing: positive [n] values for AdaptiveZone (flt)")
	n := AdaptiveZoneFlt(0.99, 1.16724, 80.0, 20.0)
	if n != 0.003168796813511676 {
		t.Error("AdaptiveZoneFlt() not expected value of 0.(00): " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}
