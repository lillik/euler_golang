package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Problem 6");
	t0 := time.Now()

	var sumSquare int = 0;
	var squareSum int = 0;

	for i := 1; i <= 100; i++ {
		sumSquare += i*i;
		squareSum += i;
	}
	result := (squareSum * squareSum) - sumSquare;
	fmt.Println("The result is:",result);

	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
