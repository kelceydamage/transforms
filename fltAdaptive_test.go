package transforms

import (
	"fmt"
	"strconv"
	"testing"
)

// FltSlice tests with negative integers
//-------------------------------------------------------------------------------------------------- <-100

func ExampleAdaptiveZoneFlt() {
	var processed []float64
	signal := []float64{0.1, 0.23, 0.44, 0.25, 0.33, 0.72, 0.48, 0.26, 0.15, 0.3}
	for _, v := range signal {
		processed = append(processed, AdaptiveZoneFlt(v, 1.16724, 80.0, 20.0))
	}
	fmt.Println(processed)
}

func TestAdaptiveZoneFlt(t *testing.T) {
	t.Log("Testing: positive [n] values for AdaptiveZone (flt)")
	n := AdaptiveZoneFlt(0.99, 1.16724, 80.0, 20.0)
	if n != 0.003168796813511676 {
		t.Error("AdaptiveZoneFlt() not expected value of 0.(00): " + strconv.FormatFloat(n, 'f', -1, 64))
	}
}
