package main

import (
	"fmt"
	"os"

	tm "github.com/buger/goterm"
	au "github.com/logrusorgru/aurora"
)

var verbal = false
var visual = true
var steps int

var puzzle = [9][9]int{
	{8, -1, 2, -1, -1, 6, -1, 5, -1},
	//{-1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, 4, -1, -1, 1, 8, -1, -1, -1},
	{-1, 9, -1, -1, -1, 3, -1, 8, 4},
	{2, -1, -1, -1, -1, 9, 8, -1, 1},
	{-1, 1, -1, -1, -1, -1, 5, 4, 9},
	//		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
	{-1, 8, -1, -1, 3, -1, 6, -1, -1},
	{-1, 7, 8, 9, -1, 2, 4, -1, 5},
	{-1, 2, 9, -1, -1, 5, 7, -1, 3},
	//{-1, -1, -1, -1, -1, -1, -1, -1, -1},
	{5, -1, 1, -1, 7, -1, 9, -1, 8}}

func main() {

	if visual {
		printBoard(puzzle, 1, 1)
	}
	solve(puzzle)

}
func solve(p [9][9]int) {
	r, c := findNextEmpty(p)
	if verbal {
		fmt.Println(r, c)
	}
	if r == -1 {
		if c == -1 {

			if visual {
				printBoard(p, 9, 9)
			}
			fmt.Println("puzzle solved in", au.Red(steps), "steps")

			os.Exit(0)
		}
	}
	if verbal {
		fmt.Println("starting search for", "row", r, "column", c)
	}
	for i := 1; i <= 9; i++ {
		if verbal {
			fmt.Println("trying", i, "at", r, ":", c)
		}
		if isValid(p, i, r, c) {
			if verbal {
				fmt.Println(i, "looks like it might work at ", r, ":", c)
			}
			p[r][c] = i
			//time.Sleep(1 * time.Duration(50000000))
			steps++
			if visual {
				printBoard(p, r, c)
			}
			solve(p)
		}

	}

}

func findNextEmpty(p [9][9]int) (x, y int) {
	for i, r := range p {
		for j, c := range r {
			if c == -1 {
				if verbal {
					fmt.Println("found", i, j, "to be empty")
				}
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(p [9][9]int, g, r, c int) bool {
	row_vals := p[r]

	if intInSlice(row_vals, g) {
		return false
	}

	col_vals := getPuzzleColumn(p, c)
	if intInSlice(col_vals, g) {
		return false
	}

	quad_vals := getPuzzleQuad(p, r, c)
	if intInSlice(quad_vals, g) {
		return false
	}

	return true

}

func intInSlice(s [9]int, v int) bool {
	for _, lv := range s {
		if lv == v {
			return true
		}
	}
	return false
}

func getPuzzleColumn(p [9][9]int, c int) [9]int {
	var o [9]int
	for x, y := range p {

		o[x] = y[c]
	}
	return o
}

func getPuzzleQuad(p [9][9]int, r, c int) [9]int {
	var o [9]int

	for i := int64(r/3) * 3; i <= int64(r/3)*3+2; i++ {
		for j := int64(c/3) * 3; j <= int64(c/3)*3+2; j++ {
			for x, y := range o {
				if y == 0 {
					o[x] = p[i][j]
					break
				}
			}
		}
	}
	return o
}

func printBoard(board [9][9]int, hx, hy int) {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Printf("=====================\n")
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			o := board[y][x]
			if o == -1 {
				o = 0
			}
			if x == hx && y == hy {
				tm.Printf("%d ", au.Green(o))
			} else {
				tm.Printf("%d ", o)
			}
			if x%3 == 2 && x < 8 {
				tm.Printf("| ")
			}
		}
		tm.Printf("\n")
		if y%3 == 2 && y < 8 {
			tm.Printf("-----   -----   -----\n")
		}
	}
	tm.Printf("=====================\n\n")
	tm.Println("Step:", au.Red(steps))
	tm.Flush()

}

/*
puzzle := [10][10]int{{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1}}
*/
