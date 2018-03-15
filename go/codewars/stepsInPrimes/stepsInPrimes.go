package stepsInPrimes

import "math"
import "fmt"

func isPrime(value int) bool {
    for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func findNextPrime(p int, max int, distance int) int{
	fmt.Printf("Finding next prime of %d\n", p)
	for i:= p +1; p<=max; i++{
		if isPrime(i) && (p-i) ==distance {
			fmt.Printf("next prime of %d is %d\n", p, i)
			return i
		}
	}
	return 0
}

func Step(g, m, n int) []int {
    for i:= m; i<=n; i++{
		fmt.Printf("%d\n", i)
		if isPrime(i){
			next := findNextPrime(i, n, g)
			if next != 0 {
				return []int{i, next}
			}
		}
	}
	return nil
}
