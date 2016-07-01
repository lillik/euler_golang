package main

import (
	"testing"
)


func TestGetGreatestProduct(t *testing.T) {
	expected := 70600674;
	prime := getGreatestProduct()
	if int(prime) != expected {
		t.Errorf("getGreatestProduct = %d, MUST_BE: %d", prime, expected)
	}
}
