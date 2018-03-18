package stepsInPrimes

import "math"


func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func findNextPrime(p int, max int, distance int) int{
	for i:= p+1; i<=max && (i-p) <= distance; i++{
		if isPrime(i) && (i-p) ==distance {
			return i
		}
	}
	return 0
}

func Step(g, m, n int) []int {
    for i:= m; i<=n; i++{
		if isPrime(i){
			next := findNextPrime(i, n, g)
			if next != 0 {
				return []int{i, next}
			}
		}
	}
	return nil
}
