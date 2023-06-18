package main

import (
	"fmt"
	"math/rand"
)

const (
	kMine = 9
)

type Cell struct {
	value   int
	visible bool
}

type Matrix struct {
	rows    int
	columns int
	data    []Cell
}

func NewMatrix(rows, columns int) Matrix {
	return Matrix{
		rows:    rows,
		columns: columns,
		data:    make([]Cell, rows*columns),
	}
}

func (matrix *Matrix) at(row, column int) *Cell {
	return &matrix.data[row*matrix.columns+column]
}

type MineField struct {
	field Matrix
}

func NewMineField(rows, columns, mines int) *MineField {
	mineField := &MineField{
		field: NewMatrix(rows, columns),
	}
	mineField.initMineField(rows, columns, mines)
	return mineField
}

func (mineField *MineField) min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (mineField *MineField) max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (mineField *MineField) swap(i, j int) {
	iVal := mineField.field.at(i/mineField.field.columns, i%mineField.field.columns)
	jVal := mineField.field.at(j/mineField.field.columns, j%mineField.field.columns)

	temp := *iVal
	*iVal = *jVal
	*jVal = temp
}

func (mineField *MineField) initMineField(rows, columns, mines int) {
	sum := rows * columns
	if mines > sum {
		mines = sum
	}
	for i := 0; i < mines; i++ {
		if i < mines {
			mineField.field.at(i/columns, i%columns).value = kMine
		}
	}

	for i := 0; i < mineField.min(mines, sum-1); i++ {
		j := i + (rand.Int() % (sum - i))
		mineField.swap(i, j)
	}

	mineField.print(true)

	for row := 0; row < mineField.field.rows; row++ {
		for column := 0; column < mineField.field.columns; column++ {
			if mineField.field.at(row, column).value != kMine {
				continue
			}

			fmt.Printf("Row Column - (%d,%d)\n", row, column)

			for i := mineField.max(0, row-1); i <= mineField.min(row-1, row+1); i++ {
				for j := mineField.max(0, column-1); j <= mineField.min(column-1, column+1); j++ {
					if (i != row || j != column) && mineField.field.at(i, j).value != kMine {
						fmt.Printf("i, j - (%d,%d)\n", i, j)
						mineField.field.at(i, j).value++
					}
				}
			}
		}
	}
}

func (mineField *MineField) onClick(row, column int) bool {
	if row < 0 || row >= mineField.field.rows || column < 0 || column >= mineField.field.columns {
		return false
	}

	if mineField.field.at(row, column).visible {
		return false
	}
	mineField.field.at(row, column).visible = true
	if mineField.field.at(row, column).value == kMine {
		fmt.Println("BOOM!!!")
		return true
	}
	if mineField.field.at(row, column).value != 0 {
		return false
	}

	mineField.onClick(row-1, column)
	mineField.onClick(row+1, column)
	mineField.onClick(row, column-1)
	mineField.onClick(row, column+1)
	return false
}

func (mineField *MineField) print(showHidden bool) {
	for i := 0; i < mineField.field.rows; i++ {
		for j := 0; j < mineField.field.columns; j++ {
			if mineField.field.at(i, j).visible || showHidden {
				fmt.Printf("%d ", mineField.field.at(i, j).value)
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	mineField := NewMineField(8, 11, 7)

	fmt.Println("############## Initial State ##############")
	mineField.print(true)
	mineField.onClick(5, 1)
	mineField.print(false)
	mineField.onClick(2, 6)
	mineField.print(false)
	mineField.onClick(9, 3)
	mineField.print(false)
	mineField.onClick(0, 0)
	mineField.print(false)
	mineField.onClick(3, 5)
	mineField.print(false)
}
