package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Problem 3");
	t0 := time.Now()
	result := getPrimeFactorsToNumber(600851475143, generatePrimeNumbers(10000));

	fmt.Println("The result is:", maxIntSlice(result));
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

//
func generatePrimeNumbers(limit int) []int {
	pp := 2;
	primList := []int{pp};

	for pp < limit {
		pp ++;
		for _, a := range primList {
			if pp % a == 0 {
				break;
			} else if in_array(pp, primList) == false {
				primList = append(primList, pp);
			}

		}

	}
	return primList;
}

//
func getPrimeFactorsToNumber (n int, primes []int) []int {
	result := []int{};
	for _, prime :=range primes {
		if n % prime == 0 {
			result = append(result, prime);
			result = append(result, getPrimeFactorsToNumber((n/prime), primes)...);
			break;
		}
	}
	return  result;
}

// check if array has element
func in_array(e int, s []int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func maxIntSlice(s []int) (max int) {

	max =0;

	for _, item := range s {
		if max < item {
			max = item;
		}
	}
	return;
}

