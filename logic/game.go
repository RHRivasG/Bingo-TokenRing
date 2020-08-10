package logic

import (
	"bingo-tokenring/logic/factories"
	"bingo-tokenring/logic/items"
	"bingo-tokenring/protocol"
	"log"
	"strings"
	"time"

	"github.com/labstack/echo"
)

//Game .
type Game struct {
	Protocol protocol.Protocol
	Blower   items.Blower
	Boards   []items.Board
	Mode     string
	Director bool
	Message  items.Message
}

//GUI .
type GUI struct {
	Ball   items.Ball    `json:"ball"`
	Boards []items.Board `json:"boards"`
	Bingo  []string      `json:"bingo"`
}

//ABC .
var ABC string = "abcdefghijklmnopqrstuvwxyz"

//Message .
var Message items.Message

//NewGame .
func NewGame(listener string, writer string, numBoards int, mode string) Game {
	var game Game
	game.Protocol, game.Boards, game.Blower = factories.BingoFactory(listener, writer, numBoards)
	game.Mode = mode
	game.Director = false
	return game
}

//LoadGame .
func (g *Game) LoadGame(ctx echo.Context) error {
	//Director
	var message []string
	message = append(message, g.Protocol.GetWriterName())
	res, err := g.ContactPlayer(message)
	if err != nil {
		log.Fatal(err)
	}
	if res[0] == g.Protocol.GetWriterName() {
		g.Director = true
		g.WriteToPlayer(message)
	} else {
		g.WriteToPlayer(res)
	}

	//Board Name
	var name []string
	name = append(name, "")
	if g.Director {
		time.Sleep(3 * time.Second)
		for i := 0; i < len(g.Boards); i++ {
			name[0] = name[0] + string([]rune(ABC)[i])
		}
		g.SetBoardName(name[0])
		res, err = g.ContactPlayer(name)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		res, err = g.ListenToPlayer()
		if err != nil {
			log.Fatal(err)
		}
		req := g.SetBoardName(res[0])
		g.WriteToPlayer(req)
	}

	//GUI
	var gui GUI
	gui.Boards = g.Boards

	return ctx.JSON(200, gui)
}

//ListenToPlayer .
func (g *Game) ListenToPlayer() ([]string, error) {
	return g.Protocol.Listen()
}

//WriteToPlayer .
func (g *Game) WriteToPlayer(message []string) {
	g.Protocol.Write(message)
}

//ContactPlayer .
func (g *Game) ContactPlayer(message []string) ([]string, error) {
	res, err := g.Protocol.Converse(message)
	return res, err
}

//SetBoardName .
func (g *Game) SetBoardName(nameTaken string) []string {
	var name []string
	strings.Trim(ABC, nameTaken)
	name = append(name, nameTaken)
	for i, b := range g.Boards {
		b.Name = string([]rune(ABC)[i])
		name[0] = name[0] + string([]rune(ABC)[i])
	}
	return name
}

//Init .
func (g *Game) Init() {
	if !g.Director {
		g.Wait()
	} else {
		time.Sleep(3 * time.Second)
	}
}

//Update .
func (g *Game) Update(ctx echo.Context) error {
	var gui GUI
	if g.Message.Finished != "true" {
		if g.Director {
			if g.Message.Bingo != "null" {
				ball := g.Blower.GetBallOut()
				g.Play(ball)
				g.Message.SaveBall(ball)
			} else {
				g.Message.Finished = "true"
			}
		} else {
			g.Play(g.Message.GetMessageBall())
		}
	} else {
		gui.Bingo = g.Message.GetMessageBingo()
	}
	gui.Boards = g.Boards
	gui.Ball = g.Message.GetMessageBall()
	return ctx.JSON(200, gui)
}

//Play .
func (g *Game) Play(ball items.Ball) {
	bingo := false
	for _, board := range g.Boards {
		row, column := board.Take(ball)
		if row != 1 && column != 1 {
			if g.Mode == "lineal" {
				bingo = board.CheckBoardLine(row, column)
			} else {
				bingo = board.CheckFull()
			}
			if bingo {
				g.Message.SaveWinner(board.Name)
			}
		}
	}
}

//SaveMessage .
func (g *Game) SaveMessage(res []string) {
	g.Message.Ball = res[0]
	g.Message.Bingo = res[1]
	g.Message.Finished = res[2]
}

//Send .
func (g *Game) Send() {
	var message []string
	message[0] = g.Message.Ball
	message[1] = g.Message.Bingo
	message[2] = g.Message.Finished
	g.WriteToPlayer(message)
}

//Wait .
func (g *Game) Wait() {
	res, err := g.ListenToPlayer()
	if err != nil {
		log.Fatal(err)
	}
	g.SaveMessage(res)
}
