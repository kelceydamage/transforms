//---------------------------------------------------------------------------------------------------- <-100
// Author: Kelcey Damage
// Go: 1.10

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Doc
//---------------------------------------------------------------------------------------------------- <-100

package transforms

import (
	"math"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// AdaptiveZoneFlt converts a normalized float (percent expressed as 0.0 - 1.0) and determines if it
// exceeds the normalZone and by how much.
//
//     [threshold] is a relative point of inflection where the straight line will curve. At [scale=100] the
//     curve will always from exactly at threshold.
//
//     [scale] is a y axis weight, at 100 the inflection point is == to the threshold
//
//     [suppressionFactor] envelopes the normilization of the norma I. at 1, 0.50 == 50. At 2, 0.50 == 25.
//     This can help squash signals with too much variance.
//
//     [normalizedValue] is a signal value expressed as a float.
//
// Example of a close-ended curve (x, 1.16724, 80.0, 20.0), in a range of 0=>100 a value of 100 will equal 0,
// while all values up to 0.56 (56%) will equal 100.
func AdaptiveZoneFlt(normalizedValue float64, suppressionFactor float64, threshold float64, scale float64) float64 {
	x := normalizedValue
	n := suppressionFactor
	h := threshold
	z := scale
	k := 100000.0

	// Calculate zero point
	aFunc := func(h, k float64) float64 {
		return (math.Pow(h, 2) - h + k) / math.Pow(h, 3.0)
	}

	// Calculate envelope coefficient
	bFunc := func(h, a, k float64) float64 {
		return (k - ((a * math.Pow(h, 3.0)) - (0.01 * (a * math.Pow(h, 2.0))))) / h
	}

	// Suppress x value range
	xFunc := func(x, n float64) float64 {
		return (x / n) * 100.0
	}

	a := aFunc(h, k)
	b := bFunc(h, a, k)
	x = xFunc(x, n)
	y := 100.0 - ((a*math.Pow(x, 3.0))-(0.01*a*math.Pow(x, 2.0))+b*x)/1000.0 + z
	if y >= 100.0 {
		y = 100.0
	}
	math.Round(y)
	return y
}
