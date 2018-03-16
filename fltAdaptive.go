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
// This file is to pretty up the godoc.

// Package transforms is a library of types and methods for running transform algorithms against certain
// Types
package transforms

import (
	"math"
)

// Code
//---------------------------------------------------------------------------------------------------- <-100

// AdaptiveZoneFlt ...
func AdaptiveZoneFlt(normalIndicator float64, multi float64, threshold float64, scale float64) float64 {
	x := normalIndicator
	n := multi
	h := threshold
	z := scale
	k := 100000.0
	aFunc := func(h, k float64) float64 {
		return (math.Pow(h, 2) - h + k) / math.Pow(h, 3.0)
	}

	bFunc := func(h, a, k float64) float64 {
		return (k - ((a * math.Pow(h, 3.0)) - (0.01 * (a * math.Pow(h, 2.0))))) / h
	}

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
