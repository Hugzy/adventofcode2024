package day4

import (
	"adventofcode2024/util"
	"fmt"
	"strings"
)

var (
	partOneInput = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	partTwoInput = ``
)

type wordsearch [][]string

func PartOneTest() {
	PartOne(partOneInput)
}

const XMAS = "XMAS"

func (ws wordsearch) horizontal() int {
	count := 0
	for _, row := range ws {
		r := strings.Join(row, "")
		count += strings.Count(r, XMAS)
	}

	return count
}

func (ws wordsearch) horizontalBackwards() int {
	count := 0
	for _, row := range ws {
		rowBackwards := ""
		for i := len(row) - 1; i >= 0; i-- {
			rowBackwards += row[i]
		}

		r := strings.Join(row, "")
		count += strings.Count(r, XMAS)
	}

	return count
}

func (ws wordsearch) vertical() int {
	count := 0
	for i := 0; i < len(ws[0]); i++ {
		coloumn := ""
		for _, row := range ws {
			coloumn += row[i]
		}
		count += strings.Count(coloumn, XMAS)
	}
	return count
}

func (ws wordsearch) verticalBackwards() int {
	count := 0
	for i := 0; i < len(ws[0]); i++ {
		coloumn := ""
		for j := len(ws) - 1; j >= 0; j-- {
			coloumn += ws[j][i]
		}
		count += strings.Count(coloumn, XMAS)
	}
	return count
}

func (ws wordsearch) diagonal() int {
	count := 0
	dim := len(ws[0])

	// First half of the diagonal
	for k := 0; k < dim; k++ {
		diagonal := ""
		for j := 0; j <= k; j++ {
			i := k - j
			diagonal += ws[i][j]
		}

		count += strings.Count(diagonal, XMAS)
		count += strings.Count(Reverse(diagonal), XMAS)
	}

	// second half of the diagonal
	for k := dim - 2; k >= 0; k-- {
		diagonal := ""
		for j := 0; j <= k; j++ {
			i := k - j
			diagonal += ws[dim-j-1][dim-i-1]
		}

		count += strings.Count(diagonal, XMAS)
		count += strings.Count(Reverse(diagonal), XMAS)
	}

	return count
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PartOne(input string) {
	matrix := make(wordsearch, 0)
	row := strings.Split(input, util.N)
	for _, v := range row {
		c := strings.Split(v, "")
		fmt.Println(c)
		matrix = append(matrix, c)
	}

	// fmt.Println(input)

	count := matrix.horizontal()
	count += matrix.horizontalBackwards()
	count += matrix.vertical()
	count += matrix.verticalBackwards()
	count += matrix.diagonal()
	fmt.Println(count)
}
