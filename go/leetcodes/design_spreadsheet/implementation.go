package leetcodesgo

import (
	"strconv"
	"strings"
	"unicode"
)

type Spreadsheet struct {
	data map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{make(map[string]int, 26*rows)}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.data[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	delete(this.data, cell)
}

func (this *Spreadsheet) parseFormulaCell(input string) int {
	if unicode.IsLetter(rune(input[0])) {
		return this.data[input]
	}
	value, _ := strconv.Atoi(input)
	return value
}

func (this *Spreadsheet) GetValue(formula string) int {
	plusIndex := strings.IndexByte(formula, '+')
	return this.parseFormulaCell(formula[1:plusIndex]) + this.parseFormulaCell(formula[plusIndex+1:])
}
