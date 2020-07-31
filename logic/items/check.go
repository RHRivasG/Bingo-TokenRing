package items

import "fmt"

//CheckColumn .
func (board *Board) CheckColumn(column int, row int) bool {
	if row != -1 {
		if board.Tiles[row][column].Taken {
			fmt.Println(row, column, board.Tiles[row][column].Taken)
			return board.CheckColumn(column, row-1)
		}
		return false
	}
	return true
}

//CheckRow .
func (board *Board) CheckRow(column int, row int) bool {
	if column != -1 {
		if board.Tiles[row][column].Taken {
			return board.CheckRow(column-1, row)
		}
		return false
	}
	return true
}

//CheckDiagonalPrincipal .
func (board *Board) CheckDiagonalPrincipal(column int, row int) bool {
	if column != -1 && row != -1 {
		if board.Tiles[row][column].Taken {
			return board.CheckDiagonalPrincipal(column-1, row-1)
		}
		return false
	}
	return true
}

//CheckDiagonalSecondary .
func (board *Board) CheckDiagonalSecondary(column int, row int) bool {
	if row != -1 {
		if board.Tiles[row][column].Taken {
			return board.CheckDiagonalSecondary(column+1, row-1)
		}
		return false
	}
	return true
}

//CheckSpecial .
func (board *Board) CheckSpecial() bool {
	if board.Tiles[0][0].Taken &&
		board.Tiles[0][4].Taken &&
		board.Tiles[4][0].Taken &&
		board.Tiles[4][4].Taken &&
		board.Tiles[2][2].Taken {
		return true
	}
	return false
}

//CheckFull .
func (board *Board) CheckFull() bool {
	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if board.Tiles[row][column].Taken == false {
				return false
			}
		}
	}
	return true
}
