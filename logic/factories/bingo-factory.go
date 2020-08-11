package factories

import (
	"bingo-tokenring/logic/items"
	"bingo-tokenring/protocol"
	"log"
)

//BingoFactory .
func BingoFactory(listener string, writer string, numBoards int) (protocol.Protocol, []items.Board, items.Blower) {
	protocol, err := protocol.NewProtocol(listener, writer)
	if err != nil {
		log.Fatal(err)
	}
	var boards []items.Board
	for i := 0; i < numBoards; i++ {
		boards = append(boards, items.NewBoard(""))
	}
	blower := items.NewBlower()
	return protocol, boards, blower

}
