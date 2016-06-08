/**
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

import (
	"testing"
)

var expectedResults = map[uint64]uint64{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
	20: 6765,
}

// func resetLookupTable() {
// 	lookupTable = map[uint64]unit64{}
// }

func TestFibonacci(t *testing.T) {
	var i uint64
	for itemCount := uint64(len(expectedResults)); i < itemCount; i++ {
		if result := Fibonacci(i); result != expectedResults[i] {
			t.Errorf("Fibonacci(%d) returned %d, expected %d", i, result, expectedResults[i])
		}
	}
}
