package gapInPrimes

import (
	"math"
)


func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func findNextPrime(p int, max int) (int, int) {
	for i:= p+1; i<=max; i++{
		if isPrime(i) {
			return i, i-p
		}
	}
	return 0,0
}


/*
Suboptimal, it should start at the first prime, not in n
*/
func Gap(g, m, n int) []int {
	i:=m
	for  i<=n {
		if isPrime(i) {
			np, d :=findNextPrime(i, n)
			if d == 0 {
				return nil
			}else {
				if d==g {
					return []int{i, np}
				} else {
					i=np
				}
			}
		}else{
			i++
		}
	}
	return nil
}

