package transforms

import (
	"strconv"
	"testing"
)

// FltSlice tests with negative integers
//-------------------------------------------------------------------------------------------------- <-100

func ExampleAdaptiveZoneInt() {
	AdaptiveZoneInt(99, 1.16724*100, 80.0, 20.0)
}

func TestAdaptiveZoneInt(t *testing.T) {
	t.Log("Testing: positive [n] values for AdaptiveZone (int)")
	n := AdaptiveZoneInt(99, 1.16724*100, 80.0, 20.0)
	if n != 0 {
		t.Error("AdaptiveZoneFlt() not expected value of 0: " + strconv.Itoa(n))
	}
}
