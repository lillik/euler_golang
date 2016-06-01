package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Problem 5");
	t0 := time.Now()
	var number int = 0;
	var found bool = true;
	for i:=2520; ; i += 20 {
		found = true;
		for j:=2; j<=20; j++ {
			if i % j !=0 {
				found = false;
			}
		}
		if found == true {
			number = i;
			break;
		}

	}
	fmt.Println("The number is:",number);
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
