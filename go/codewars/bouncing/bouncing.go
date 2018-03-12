package bouncing;
import (
	"math"
)


func BouncingBall(h, bounce, window float64) int {
	if (h<0 || bounce <= 0 || bounce >=1 || window>=h || window <=0){
		return -1
	} else {
		return int(math.Floor(math.Log(window/h)/math.Log(bounce))*2 + 1)
	}
}

func BouncingBall_2(h, bounce, window float64) int {
	if (h<0 || bounce <= 0 || bounce >=1 || window>=h || window <=0){
		return -1
	} else {
		count := 0
		hh := h
		for hh >= window {
			count += 1
			hh = hh * bounce
			if hh >= window {
				count += 1
			}
		}
		return count
	}
}
