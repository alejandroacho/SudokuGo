package sudoku

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func Test_if_return_a_number_greater_or_equal_than_given(test *testing.T){
	for i:=0; i<25; i++{
		result:=randomInteger(1,9)
		assert.GreaterOrEqual(test, result, 1)
	}
}
func Test_if_return_a_number_lower_or_equal_than_given(test *testing.T){
	for i:=0; i<25; i++{
		result:=randomInteger(1,9)
		assert.LessOrEqual(test, result, 9)
	}
}
func Test_if_return_a_number(test *testing.T){
	result:=randomInput()
	var number int=0
	assert.IsType(test, number, result)
}
func Test_if_return_a_number_greater_or_equal_than_1(test *testing.T){
	for i:=0; i<25; i++{
		result:=randomInput()
		assert.GreaterOrEqual(test, result, 1)
	}
}
func Test_if_return_a_number_lower_or_equal_than_9(test *testing.T){
	for i:=0; i<25; i++{
		result:=randomInput()
		assert.LessOrEqual(test, result, 9)
	}
}
func Test_if_return_correct_quadrant_given_a_coordenada(test *testing.T){
	result:=quadrantPicker(1,1)
	expected:=1
	assert.Equal(test, result, expected)
}
func Test_if_return_correct_quadrant_given_a_coordenada2(test *testing.T){
	result:=quadrantPicker(5,3)
	expected:=5
	assert.Equal(test, result, expected)
}
func Test_if_return_is_an_array(test *testing.T){
	var given [9][9] int
	result:=sudokuMaker(given)
	assert.IsType(test, given, result)
}
func Test_algorithm_make_a_good_sudoku(test *testing.T){
	empty:=[9][9]int{}
	board:=sudokuMaker(empty)
	result:=verifyMatch(board)
	assert.True(test,result)
}