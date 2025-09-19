package leetcodesgo

import (
	"strconv"
	"strings"
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
	output, err := strconv.Atoi(input)
	if err != nil {
		cell := this.data[input]
		return cell
	}
	return output
}

func (this *Spreadsheet) GetValue(formula string) int {
	values := strings.Split(formula[1:], "+")
	return this.parseFormulaCell(values[0]) + this.parseFormulaCell(values[1])
}
