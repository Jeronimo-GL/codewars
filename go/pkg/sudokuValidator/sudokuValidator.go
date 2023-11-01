package sudoku


/* 
ValidateSolution 
Function that validates sudoku solutions
*/
func ValidateSolution(m [][]int) bool {
	//validate rows
	if hasZeros(m) {
		return false;
	};
	// validate rows and cols
	for i:=0; i<9; i++ {
		if !validateRow(m, i) {
			return false
		}
		if !validateCol(m, i) {
			return false
		}
	}

	// validate blocks
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if !validateBlock(m, i, j){
				return false
			}
		}
	}
    return true
}

func hasZeros(m[][]int) bool{
	for row:=0; row<9; row ++{
		for col:=0; col<9; col ++{
			if m[row][col]==0 {
				return true
			}
		}
	}
	return false
}

func validateRow(m [][]int, row int) bool{
	for i:=1; i<10; i++ {
		if countInRow(m, row, i) != 1{
			return false
		}
	}
	return true
}

func countInRow(m [][]int, row int, value int) int {
	var count=0;
	for col:=0; col <9; col++ {
		if m[row][col] == value {
			count ++
		}
	}
	return count
}

func validateCol(m [][]int, col int) bool{
	for i:=1; i<10; i++ {
		if countInCol(m, col, i) != 1{
			return false
		}
	}
	return true
}

func countInCol(m [][]int, col int, value int) int {
	var count=0;
	for row:=0; row <9; row++ {
		if m[row][col] == value {
			count ++
		}
	}
	return count
}

func validateBlock(m [][]int, brow int, bcol int) bool{
	for i:=1; i<10; i++ {
		if countInBlock(m, brow, bcol, i) != 1{
			return false
		}
	}
	return true
}

func countInBlock(m [][]int, brow int, bcol int, value int) int {
	var count=0;
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if m[brow*3+i][bcol*3 + j] == value {
				count ++
			}
		}
	}
	return count;
}

