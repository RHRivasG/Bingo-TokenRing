package items

import (
	"math/rand"
	"strings"
	"time"
)

//Board .
type Board struct {
	Name  string     `json:"name"`
	Tiles [5][5]Tile `json:"tiles"`
}

// NewBoard .
func NewBoard(name string) Board {
	var board Board
	board.Name = name
	vingu := []string{"V", "I", "N", "G", "U"}
	for i, letter := range vingu {
		for j := 0; j < 5; j++ {
			board.Tiles[j][i] = board.SetTileRandomNumber(letter, i, j, i*15+1, i*15+16)
		}
	}
	board.Tiles[2][2].Taken = true
	return board
}

//SetTileRandomNumber .
func (board *Board) SetTileRandomNumber(letter string, column int, row int, min int, max int) Tile {
	n := Random(min, max)
	if board.SeeNumberInColumn(column, row, n) != -1 {
		return board.SetTileRandomNumber(letter, column, row, min, max)
	}
	return Tile{letter, n, false}
}

//Take .
func (board *Board) Take(ball Ball) (int, int) {
	vingu := "VINGU"
	column := strings.Index(vingu, ball.Letter)
	if column != -1 {
		if row := board.SeeNumberInColumn(column, 4, ball.Number); row != -1 {
			board.Tiles[row][column].Taken = true
			return row, column
		}
	}
	return -1, -1
}

//SeeNumberInColumn .
func (board *Board) SeeNumberInColumn(column int, row int, n int) int {
	if row != -1 {
		if board.Tiles[row][column].Number == n {
			return row
		}
		return board.SeeNumberInColumn(column, row-1, n)
	}
	return -1
}

//CheckBoardLine .
func (board *Board) CheckBoardLine(row int, column int) bool {
	if board.CheckColumn(column, 4) {
		return true
	} else if board.CheckRow(4, row) {
		return true
	} else if board.CheckDiagonalPrincipal(4, 4) {
		return true
	} else if board.CheckDiagonalSecondary(0, 4) {
		return true
	} else if board.CheckSpecial() {
		return true
	}
	return false
}

//Random min <= n < max
func Random(min int, max int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(max - min)
	return n + min
}
