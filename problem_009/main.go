package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Problem 9");
	t0 := time.Now()

	result := findPythagoreanTripletBySum(1500);

	fmt.Println("The result is:", result);
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

//
func findPythagoreanTripletBySum(sum int) int {
	result := 0;
	for i:=1; i<= sum-2; i++ {
		for j:=i+1; j<= sum-1; j++ {
			c := sum-i-j;
			if (c < j) {
				break;
			}
			if (i*i + j*j == c*c) {
				result =  i*j*c;
				return result;
			}
		}
	}
	return result;
}