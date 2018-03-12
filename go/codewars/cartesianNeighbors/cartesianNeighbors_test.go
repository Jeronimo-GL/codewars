package cartesianNeighbors

import "testing"

func equalsArray(a,b []int) bool {
	return (a[0]==b[0] && a[1]==b[1]) || (a[0]==b[1] && a[1]==b[0])
}

func arrayContains(list [][]int, v []int) bool {
	for _, n := range list {
		if equalsArray(v,n) {
			return true
		}
	}
	return false
}


func Test_provided1(t *testing.T){
	point := []int{2,2}
	list := [][]int {[]int{1,1},[]int{1,2},[]int{1,3},[]int{2,1},[]int{2,3},[]int{3,1},[]int{3,2},[]int{3,3}}
	resp:= CartesianNeighbor(point[0], point[1])
	
	for _, n := range resp {
		if !arrayContains(list, n) {
			t.Errorf("(%d, %d) is not in the expected list", n[0], n[1])
		}
	}
}

func Test_provided2(t *testing.T){
	point := []int{5,7}
	list := [][]int {[]int{6,7},[]int{6,6},[]int{6,8},[]int{4,7},[]int{4,6},[]int{4,8},[]int{5,6},[]int{5,8}}
	resp:= CartesianNeighbor(point[0], point[1])
	
	for _, n := range resp {
		if !arrayContains(list, n) {
			t.Errorf("(%d, %d) is not in the expected list", n[0], n[1])
		}
	}
}

func Test_zero(t *testing.T){
	point := []int{0,0}
	list := [][]int {[]int{-1,1},[]int{0,1},[]int{1,1},[]int{-1,0},[]int{1,0},[]int{-1,-1},[]int{0,-1},[]int{1,-1}}
	resp:= CartesianNeighbor(point[0], point[1])
	
	for _, n := range resp {
		if !arrayContains(list, n) {
			t.Errorf("(%d, %d) is not in the expected list", n[0], n[1])
		}
	}
}

