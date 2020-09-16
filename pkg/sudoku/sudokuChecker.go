package sudoku

func verifyAllRows(board [9][9]int)bool{
	for x:=0; x<9; x++ {
		row:=board[x]
		contador:=0
		for y:=0; y<9; y++ {
			contador=contador+row[y]
		}
		if contador>45{
			return false
		}
		if contador<45{
			return false
		}
	}
	return true
}

func verifyAllColumns(board [9][9] int)bool{

	for x:=0; x<9; x++{
		columna:=transformColumnToArray(board,x)
		contador:=0
		for y:=0; y<9; y++ {
			contador=contador+columna[y]
		}
		if contador>45{
			return false
		}
		if contador<45{
			return false
		}
	}
	return true
}

func verifyAllQuadrants(board [9][9] int)bool{
	for q:=1; q<10; q++{
		quadrant:=quadrantDefiner(board, q)
		contador:=0
		for x := 0; x<3; x++ {
			for y := 0; y<3; y++ {
					contador=contador+quadrant[x][y]
			}
		}
		if contador!=45{
			return false
		}
	}
	return true
}

func verifyMatch(board[9][9]int) bool{
	if verifyAllColumns(board)==verifyAllRows(board)==verifyAllQuadrants(board)==true{
		return true
	}
	return false
}