package main

import (
	"testing"
)


type testpair struct {
	number   int
	expected int
}

var tests = []testpair{
	{1000, 31875000},
	{1500, 97500000},
}


func TestFindPythagoreanTripletBySum(t *testing.T) {
	for _, test := range tests {
		prime := findPythagoreanTripletBySum(test.number)
		if prime != test.expected {
			t.Errorf("sumPrimeNumber(%d) = %d, MUST_BE: %d", test.number, prime, test.expected)
		}
	}
}
