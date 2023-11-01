package cartesianDistance

import "math"


func CartesianNeighborsDistance(x, y, r int) ([]float64){
	var resp []float64
	
	for i:=1; i<=r ; i++{
		for j:=0; j<=i ; j++ {
			d:=math.Sqrt(float64(i*i) + float64(j*j))
			if !arrayContains(resp, d) {
				resp=append(resp, d)
			}
		}
	}

  return resp
}

func arrayContains(list []float64, v float64) bool {
	for _, n := range list {
		if v==n {
			return true
		}
	}
	return false
}
