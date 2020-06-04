package sudoku

import (
    "math/rand"
	"time"
	"fmt"
)
func SudokuGenerator()[9][9]int{
	board:=[9][9]int{}
	sudoku:=sudokuMaker(board)
	hidden:=hideNumbers(sudoku)
	return hidden
}

func randomInteger(min, max int) int {
    return min + rand.Intn(max-min)
}

func randomInput()int{
	rand.Seed(time.Now().UnixNano())
	randomNumber:=randomInteger(1, 10)
	return randomNumber
}

func quadrantPicker(x int, y int) int{
	var quadrantNumber int
	if x<3 && y<3{ quadrantNumber=1 }
	if x<3 && y>=3 && y<6{ quadrantNumber=2 }
	if x<3 && y>=6 && y<9{ quadrantNumber=3 }
	if x>=3 && x<6 && y<3{ quadrantNumber=4 }
	if x>=3 && x<6 && y>=3 && y<6{ quadrantNumber=5 }
	if x>=3 && x<6 && y>=6 && y<9{ quadrantNumber=6 }
	if x>=6 && x<9 && y<3{ quadrantNumber=7 }
	if x>=6 && x<9 && y>=3 && y<6{ quadrantNumber=8 }
	if x>=6 && x<9 && y>=6 && y<9{ quadrantNumber=9 }
	return quadrantNumber
}

func sudokuMaker(board[9][9]int)[9][9]int{
	sudoku:=board
	for x:=0; x<9; x++{
		for y:=0; y<9; y++{		
			contador:=0
			contadorUltra:=0			
			for{
				if contador==2000{
					sudoku[x]=[9]int{0,0,0,0,0,0,0,0,0}
					contador=0
					y=0
				}
				if contadorUltra==4000{
					sudoku[x]=[9]int{0,0,0,0,0,0,0,0,0}
					sudoku[x-1]=[9]int{0,0,0,0,0,0,0,0,0}
					contadorUltra=0
					contador=0
					y=0
					x=x-1
				}
				number:=randomInput()
				row:=numberInRow(sudoku[x],number)
				col:=numberInColumn(transformColumnToArray(sudoku,y),number)
				quadrantNumber:=quadrantPicker(x, y)
				quadrant:=numberInQuadrant(quadrantDefiner(sudoku,quadrantNumber),number)
				superTrue:=verifyAllInputBooleans(row, col, quadrant)
				if superTrue==true{
					sudoku=makeInput(sudoku, x, y, number)
					/*fmt.Println(sudoku)*/
					break
				}
				if superTrue==false{
					contadorUltra++
					contador++
				}
			}
		}
	}
	fmt.Println(sudoku)
	hidden:=hideNumbers(sudoku)
	fmt.Println(hidden)
	fmt.Println()
	return sudoku
}

func hideNumbers(sudoku[9][9]int)[9][9]int{
	for i:=0; i<70; i++{
		x:=randomInteger(0,9)
		y:=randomInteger(0,9)
		sudoku[x][y]=0
	}
	return sudoku
}

