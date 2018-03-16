package transforms

import (
	"fmt"
	"strconv"
	"testing"
)

// FltSlice tests with negative integers
//-------------------------------------------------------------------------------------------------- <-100

func ExampleAdaptiveZoneInt() {
	var processed []int
	signal := []int{1, 23, 44, 25, 33, 72, 48, 26, 15, 3}
	for _, v := range signal {
		processed = append(processed, AdaptiveZoneInt(v, 1.16724*100, 80.0, 20.0))
	}
	fmt.Println(processed)
}

func TestAdaptiveZoneInt(t *testing.T) {
	t.Log("Testing: positive [n] values for AdaptiveZone (int)")
	n := AdaptiveZoneInt(99, 1.16724*100, 80.0, 20.0)
	if n != 0 {
		t.Error("AdaptiveZoneFlt() not expected value of 0: " + strconv.Itoa(n))
	}
}
