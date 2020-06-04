package sudoku

func makeInput(board[9][9] int, fila int, columna int, number int) [9][9]int{
	if number>0 && number<=9{
		board[fila][columna]=number
		return board
	}
	return board
}

func numberInRow(row [9] int, number int)bool{	
	for i:= 0; i<9; i++{
		if number == row[i] {
			return false
		}
	} 
	return true                                     
}

func transformColumnToArray(board[9][9]int, numberCol int)[9]int{
	column:=[9] int {}
	for x:=0; x<9;x++{
		column[x]=board[x][numberCol]
	}
	return column
}

func numberInColumn(col [9] int, number int)bool{
	for i:= 0; i<9; i++{
		if number == col[i] {
			return false
		}
	} 
	return true
}

func quadrantDefiner(board[9][9] int, quadrantNumber int)[3][3]int{
	quadrant:=[3][3]int{}
	if quadrantNumber==1{
		for x:= 0; x<3; x++{
			for y:=0; y<3; y++{
				quadrant[x][y]=board[x][y]
			}
		} 
	}
	if quadrantNumber==2{
		for x:= 0; x<3; x++{
			contador:=0
			for y:=3; y<6; y++{
				quadrant[x][contador]=board[x][y]
				contador ++
			}
		} 
	}
	if quadrantNumber==3{
		for x:= 0; x<3; x++{
			contador:=0
			for y:=6; y<9; y++{
				quadrant[x][contador]=board[x][y]
				contador ++
			}
		} 
	}
	if quadrantNumber==4{
		contadorX:=0
		for x:=3; x<6; x++{
			contadorY:=0
			for y:=0; y<3; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	if quadrantNumber==5{
		contadorX:=0
		for x:=3; x<6; x++{
			contadorY:=0
			for y:=3; y<6; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	if quadrantNumber==6{
		contadorX:=0
		for x:=3; x<6; x++{
			contadorY:=0
			for y:=6; y<9; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	if quadrantNumber==7{
		contadorX:=0
		for x:=6; x<9; x++{
			contadorY:=0
			for y:=0; y<3; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	if quadrantNumber==8{
		contadorX:=0
		for x:=6; x<9; x++{
			contadorY:=0
			for y:=3; y<6; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	if quadrantNumber==9{
		contadorX:=0
		for x:=6; x<9; x++{
			contadorY:=0
			for y:=6; y<9; y++{
				quadrant[contadorX][contadorY]=board[x][y]
				contadorY ++
			}
			contadorX++
		} 
	}
	return quadrant
}

func numberInQuadrant(quadrant [3][3]int, number int)bool{
	for x := 0; x<3; x++ {
		for y := 0; y<3; y++ {
			if number == quadrant[x][y] {
				return false
			}
				
		}
	}
	return true
}

func verifyAllInputBooleans(row bool, col bool, quadrant bool)bool{
	if row==false{
		return false
	}
	if col==false{
		return false
	}
	if quadrant==false{
		return false
	}
	return true
}