package magnetParticules

import (
	. "math"
)

func v(k float64, n float64) float64{
	return 1 /(k*Pow((n+1), 2*k))
}

func Doubles(maxk, maxn int) float64 {
	var sum float64 = 0
    for k:=1; k<=maxk; k++ {
		for n:=1; n<=maxn; n++ {
			sum += v(float64(k),float64(n))
		}
	}
	
	return sum
}
