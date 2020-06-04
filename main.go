package main

import (
	"sudoku/pkg/sudoku"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)
//Cell object
type Cell struct {
	value     int
	column    int
  row       int
  isVisible bool // true: se ve el número y false: no se muestra el núm
}
func generateSudoku() [9][9]int {
	board := sudoku.SudokuGenerator()
	return board
}

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 800,
		Title:  "Sudoku",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(generateSudoku)
	app.Run()
}
