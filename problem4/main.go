package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	fmt.Println("Problem 4");
	t0 := time.Now()
	// define largest palindrome number
	var lpn int = 0;
	//Execute Algorithm
	for i:= 999; i > 100; i-- {
		//fmt.Println(i);
		for j:= 999; j > 100; j-- {
			number := i * j;
			// break if number is < largest palindrome number
			if (lpn > number) {
				break;
			}
			// check if is palindromic number
			if (is_palindromic_number(number)) {
				lpn = number;
			}
		}
	}
	fmt.Println("The largest palindrome number is:", lpn);
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}


// Check if number is Palindromic Number
func is_palindromic_number(n int) bool {
	var s string = strconv.Itoa(n);
	var endOfIndex int = 0;
	length := len(s);

	for i := 0; i< (length/2) ; i++ {
		endOfIndex = length-1-i;
		if s[i] != s[endOfIndex] {
			return false;
		}
	}

	return true;
}
