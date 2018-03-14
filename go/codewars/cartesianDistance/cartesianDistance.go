package cartesianDistance

import "math"
import "fmt"


func CartesianNeighborsDistance(x, y, r int) ([]float64){
	var resp []float64
	
	for i:=1; i<=r ; i++{
		for j:=0; j<=i ; j++ {
			fmt.Printf("%f\n", math.Sqrt(float64(i*i) + float64(j*j)))
			resp=append(resp, math.Sqrt(float64(i*i) + float64(j*j)))
		}
	}

  return resp
}
