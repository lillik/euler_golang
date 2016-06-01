package main

import (
	"testing"
)


type testpair struct {
	number   int
	expected int
}

var tests = []testpair{
	{10, 17},
	{2000000, 142913828922},
}


func TestSumPrimeNumber(t *testing.T) {
	for _, test := range tests {
		prime := sumPrimeNumber(test.number)
		if prime != test.expected {
			t.Errorf("sumPrimeNumber(%d) = %d, MUST_BE: %d", test.number, prime, test.expected)
		}
	}
}
