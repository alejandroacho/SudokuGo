package sudoku

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_checks_return_true_if_all_rows_ok(test *testing.T){
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
	result:= verifyAllRows(board)
	assert.True(test,result)
}
func Test_checks_return_false_if_one_row_is_not_ok(test *testing.T){
	board:=[9][9]int{
		{5, 5, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	result:= verifyAllRows(board)
	assert.False(test,result)
}
func Test_checks_return_true_if_all_columns_ok(test *testing.T){
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
	result:= verifyAllColumns(board)
	assert.True(test,result)
}
func Test_checks_return_false_if_one_column_is_not_ok(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{5, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	result:= verifyAllColumns(board)
	assert.False(test,result)
}
func Test_checks_if_all_quadrants_are_ok_return_true(test *testing.T){
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
	result:= verifyAllQuadrants(board)
	assert.True(test,result)
}
func Test_checks_if_one_quadrant_isnt_ok_return_false(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 5, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	result:= verifyAllQuadrants(board)
	assert.False(test,result)
}
func Test_checks_if_a_match_is_ok_return_true(test *testing.T){
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
	result:= verifyMatch(board)
	assert.True(test,result)
}
func Test_checks_if_a_match_isnt_ok_return_false(test *testing.T){
	board:=[9][9]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 9, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}
	result:= verifyMatch(board)
	assert.False(test,result)
}
func Test_my_first_sudoku_made_by_algorithm_is_ok(test *testing.T){
	board:=[9][9]int{
		{5, 7, 1, 4, 9, 2, 6, 3, 8}, 
		{8, 2, 3, 1, 7, 6, 4, 9, 5},
		{9, 6, 4, 3, 8, 5, 7, 2, 1},
		{2, 5, 7, 9, 3, 1, 8, 6, 4},
		{1, 8, 9, 2, 6, 4, 3, 5, 7},
		{3, 4, 6, 8, 5, 7, 2, 1, 9},
		{7, 9, 2, 5, 4, 3, 1, 8, 6},
		{6, 1, 5, 7, 2, 8, 9, 4, 3},
		{4, 3, 8, 6, 1, 9, 5, 7, 2},
	}
	result:= verifyMatch(board)
	assert.True(test,result)
}
