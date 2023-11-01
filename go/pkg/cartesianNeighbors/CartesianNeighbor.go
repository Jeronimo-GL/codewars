package cartesianNeighbors

func CartesianNeighbor(x,y int) [][]int{
	var resp [][]int 
	for i:=-1; i<=1; i++ {
		for j:=-1; j<=1; j++ {
			if i != 0 || j != 0 {
				resp=append(resp, []int{x+j, y+i})
			}
		}
	}
	return resp
}


