package sudoku

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func Test_if_can_make_a_input(test *testing.T){
	board:=[9][9]int{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	fila:=0
	columna:=0
	numero:=1
	expected:=[9][9]int{
		{1,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	result:=makeInput(board,fila,columna,numero)
	assert.Equal(test,expected,result)
}
func Test_if_can_make_a_input2(test *testing.T){
	board:=[9][9]int{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	fila:=4
	columna:=5
	numero:=1
	expected:=[9][9]int{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,1,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	result:=makeInput(board,fila,columna,numero)
	assert.Equal(test,expected,result)
}
func Test_if_cant_make_a_input_of_bigger_than_9(test *testing.T){
	board:=[9][9]int{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	fila:=4
	columna:=5
	numero:=10
	expected:=[9][9]int{
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0},
	}
	result:=makeInput(board,fila,columna,numero)
	assert.Equal(test,expected,result)
}
func Test_returns_false_when_number_is_repeated_in_row(test *testing.T){
	row:=[9]int{0, 0, 0, 8, 0, 3, 0, 0, 0}
	result:= numberInRow(row, 8)
	assert.False(test,result)
}
func Test_if_number_is_not_repeated_in_row(test *testing.T){ 
	row:=[9]int{0, 0, 0, 8, 0, 3, 0, 0, 0}
	result:= numberInRow(row, 1)
	assert.True(test,result)
}

func Test_returns_false_when_number_is_repeated_in_column(test *testing.T){
	row:=[9]int{0, 0, 0, 8, 0, 3, 0, 0, 0}
	result:= numberInRow(row, 8)
	assert.False(test,result)
}
func Test_given_a_column_convert_it_to_array(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	expected := [9] int {5,6,1,8,4,7,9,2,3}
	numberCol:=0
	result:=transformColumnToArray(board, numberCol)
	assert.Equal(test,expected, result)
}

func Test_returns_true_when_number_is_not_repeated_in_column(test *testing.T){
	col:=[9]int{0, 0, 0, 0, 0, 3, 0, 0, 0}
	result:= numberInColumn(col, 8)
	assert.True(test,result)
}
func Test_returns_a_quadrant_when_board_and_number_given(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=1
	expected:= [3][3]int{
		{5,3,4},
		{6,7,2},
		{1,9,8},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_2(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=2
	expected:= [3][3]int{
		{6,7,8},
		{1,9,5},
		{3,4,2},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_3(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=3
	expected:= [3][3]int{
		{9,1,2},
		{3,4,8},
		{5,6,7},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_4(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=4
	expected:= [3][3]int{
		{8,5,9},
		{4,2,6},
		{7,1,3},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_5(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=5
	expected:= [3][3]int{
		{7,6,1},
		{8,5,3},
		{9,2,4},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_6(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=6
	expected:= [3][3]int{
		{4,2,3},
		{7,9,1},
		{8,5,6},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_7(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=7
	expected:= [3][3]int{
		{9,6,1},
		{2,8,7},
		{3,4,5},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_8(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=8
	expected:= [3][3]int{
		{5,3,7},
		{4,1,9},
		{2,8,6},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_a_quadrant_when_board_and_number_given_9(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	quadrantNumber:=9
	expected:= [3][3]int{
		{2,8,4},
		{6,3,5},
		{1,7,9},
	}
	result:= quadrantDefiner(board,quadrantNumber)
	assert.Equal(test,expected,result)
}
func Test_returns_false_when_number_is_repeated_in_quadrant(test *testing.T){
	quadrant:=[3][3]int{
		{0, 1, 2},
		{4, 0, 0},
		{0, 9, 5},
	}
	result:= numberInQuadrant(quadrant, 9)
	assert.False(test,result)
}
func Test_returns_true_when_number_is_repeated_in_quadrant(test *testing.T){
	quadrant:=[3][3]int{
		{0, 1, 2},
		{4, 0, 0},
		{0, 0, 5},
	}
	result:= numberInQuadrant(quadrant, 9)
	assert.True(test,result)
}
func Test_returns_false_when_one_is_false(test *testing.T){
	var row bool=true
	var col bool=false
	var quadrant bool=true
	result:= verifyAllInputBooleans(row, col, quadrant)
	assert.False(test,result)
}
func Test_returns_true_when_all_true(test *testing.T){
	var row bool=true
	var col bool=true
	var quadrant bool=true
	result:= verifyAllInputBooleans(row, col, quadrant)
	assert.True(test,result)
}
func Test_returns_false_when_all_false(test *testing.T){
	var row bool=false
	var col bool=false
	var quadrant bool=false
	result:= verifyAllInputBooleans(row, col, quadrant)
	assert.False(test,result)
}