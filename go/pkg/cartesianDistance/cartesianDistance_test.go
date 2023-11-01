package cartesianDistance

import "testing"

func equalsArray(a,b []int) bool {
	return (a[0]==b[0] && a[1]==b[1]) || (a[0]==b[1] && a[1]==b[0])
}


func check_response(t *testing.T, x int, y int, resp []float64, expected []float64) {
	if len(resp) != len(expected) {
		t.Errorf("Wrong response length. Expected %d, got %d", len(expected), len(resp))
	}

	for _, n := range resp {
		if !arrayContains(expected, n) {
			t.Errorf("%f is not in the expected list", n)
		}
	}
}

func Test_distance_1(t *testing.T){
	var x,y int= 1,1
	dists := []float64 {1.414214, 1.0}

	ds :=CartesianNeighborsDistance(x, y, 1)
	check_response(t, x, y,  ds, dists)
}

func Test_distance_2(t *testing.T){
	var x,y int= 1,1
	dists := []float64 {1.414214, 1.0, 2.0, 2.2360679775, 2.8284271247}

	ds :=CartesianNeighborsDistance(x, y, 2)
	check_response(t, x, y,  ds, dists)
}


func Test_distance_3(t *testing.T){
	var x,y int= 1,1
	dists := []float64 {1.414214, 1.0, 2.0, 2.2360679775, 2.8284271247, 3, 3.16227766017, 3.60555127546, 4.24264068712 }

	ds :=CartesianNeighborsDistance(x, y, 3)
	check_response(t, x, y,  ds, dists)
}
