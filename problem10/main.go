package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Problem 10");
	t0 := time.Now()

	sum := 0;
	pp :=1;
	for pp < 2000000{
		pp +=1;
		if pp == 2 || pp == 3 || pp == 5 || pp == 7 {
			sum +=pp;
		} else if (pp % 2 != 0) && (pp % 3 != 0) && (pp % 5 != 0) && (pp % 7 != 0) {
			if prime(uint32(pp)) {
				sum +=pp;

			}
		}
	}

	fmt.Println("The result is:", sum);
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}


//Using algorithm Miller–Rabin primality test.
//https://en.wikipedia.org/wiki/Miller–Rabin_primality_test
func prime(n uint32) bool {
	// bases of 2, 7, 61 are sufficient to cover 2^32
	switch n {
	case 0, 1:
		return false
	case 2, 7, 61:
		return true
	}
	// compute s, d where 2^s * d = n-1
	nm1 := n - 1
	d := nm1
	s := 0
	for d&1 == 0 {
		d >>= 1
		s++
	}
	n64 := uint64(n)
	for _, a := range []uint32{2, 7, 61} {
		// compute x := a^d % n
		x := uint64(1)
		p := uint64(a)
		for dr := d; dr > 0; dr >>= 1 {
			if dr&1 != 0 {
				x = x * p % n64
			}
			p = p * p % n64
		}
		if x == 1 || uint32(x) == nm1 {
			continue
		}
		for r := 1; ; r++ {
			if r >= s {
				return false
			}
			x = x * x % n64
			if x == 1 {
				return false
			}
			if uint32(x) == nm1 {
				break
			}
		}
	}
	return true
}
