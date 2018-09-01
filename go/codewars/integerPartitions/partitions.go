package IntegerPartitions


import (
	"math"
)

func Partitions(n int) int {
	cache := make(map[int]int)
	parts:= paux(n, cache)
	return parts
}

func paux(n int, cache map[int]int) int {
	if n <=1 {
		return 1
	} else {
		sum:= 0
		nk:= 1
		for n-sub(nk) >=0{
			part, exists:=cache[n-sub(nk)]
			if !exists {
				part=paux(n-sub(nk), cache)
				cache[n-sub(nk)]=part
			}
			sum = sum + sign(nk)*part
			nk=next(nk)
		}
		return sum
	}
}


func next(n int) int {
	if n >= 0 {
		return -n
	} else {
		return -n + 1
	}
}

func sign(k int) int {
	return int(math.Floor(math.Pow(float64(-1), math.Abs(float64(k))-1.0)))
}

func sub(k int) int{
	return k*(3*k-1)/2
}
