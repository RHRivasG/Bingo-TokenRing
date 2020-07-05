package main

import (
	item "bingo-tokenring/logic/items"
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	board := item.NewBoard()
	fmt.Println(board)
}
