package seqDisguise

import "math"

func Fct2(n uint) uint {
	if n == 0 {
		return 1
	}
	if n == 1{
		return 2
	}
	f1:=Fct(n-1)
	f2:=Fct(n-2)
    return (6*f2*f1)/(-f1+5*f2)
}

func Fct(n uint) uint {
	return uint(math.Pow(2,float64(n)))
}
