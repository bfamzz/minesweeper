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

// func (mineField *MineField) max(x, y int) int {
// 	if x > y {
// 		return x
// 	}

// 	return y
// }

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

	mineField.print(true)

	for i := 0; i < mineField.min(mines, sum-1); i++ {
		j := i + (rand.Int() % (sum - i))
		mineField.swap(i, j)
	}

	mineField.print(true)

	for row := 0; row < mineField.field.rows; row++ {
		for column := 0; column < mineField.field.columns; column++ {
			if mineField.field.at(row, column).value != kMine {
				// check the cells above, above-diagonal-left, above-diagonal-right
				// left, right
				// below, below-diagonal-left, below-diagonal-right
				rowAbove := row - 1
				leftCol := column - 1
				rightCol := column + 1
				rowBelow := row + 1
				if rowAbove >= 0 {
					if mineField.field.at(rowAbove, column).value == kMine {
						mineField.field.at(row, column).value++
					}

					if leftCol >= 0 {
						if mineField.field.at(rowAbove, leftCol).value == kMine {
							mineField.field.at(row, column).value++
						}
					}

					if rightCol < mineField.field.columns {
						if mineField.field.at(rowAbove, rightCol).value == kMine {
							mineField.field.at(row, column).value++
						}
					}
				}

				if leftCol >= 0 {
					if mineField.field.at(row, leftCol).value == kMine {
						mineField.field.at(row, column).value++
					}
				}

				if rightCol < mineField.field.columns {
					if mineField.field.at(row, rightCol).value == kMine {
						mineField.field.at(row, column).value++
					}
				}

				if rowBelow < mineField.field.rows {
					if mineField.field.at(rowBelow, column).value == kMine {
						mineField.field.at(row, column).value++
					}

					if leftCol >= 0 {
						if mineField.field.at(rowBelow, leftCol).value == kMine {
							mineField.field.at(row, column).value++
						}
					}

					if rightCol < mineField.field.columns {
						if mineField.field.at(rowBelow, rightCol).value == kMine {
							mineField.field.at(row, column).value++
						}
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
	mineField.onClick(row-1, column-1)
	mineField.onClick(row-1, column+1)

	mineField.onClick(row+1, column)
	mineField.onClick(row+1, column+1)
	mineField.onClick(row+1, column-1)

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
	rows, columns := 8, 11
	mineField := NewMineField(rows, columns, 7)

	fmt.Println("############## Initial State ##############")
	mineField.print(true)
	for {
		row, column := rand.Intn(rows), rand.Intn(columns)
		fmt.Printf("Clicked on: (%d, %d)\n", row, column)
		over := mineField.onClick(row, column)
		if over {
			break
		} else {
			mineField.print(false)
		}
	}
}
