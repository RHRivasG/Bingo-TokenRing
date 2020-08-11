package logic

import (
	"bingo-tokenring/logic/factories"
	"bingo-tokenring/logic/items"
	"bingo-tokenring/protocol"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
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

//NewGame .
func NewGame(listener string, writer string, numBoards int, mode string) Game {
	var game Game
	game.Protocol, game.Boards, game.Blower = factories.BingoFactory(listener, writer, numBoards)
	game.Mode = mode
	game.Director = false
	game.Message.Ball = "null"
	game.Message.Bingo = "null"
	game.Message.Finished = "false"
	return game
}

//LoadGame .
func (g *Game) LoadGame(ctx *gin.Context) {
	g.Protocol.Reset()
	//Director
	var message []string
	message = append(message, g.Protocol.GetWriterName())
	res, err := g.ContactPlayer(message)
	if err != nil {
		log.Fatal(err)
	}
	if res[0] == g.Protocol.GetWriterName() {
		g.Director = true
	} else {
		g.WriteToPlayer(res)
	}
	g.LoadBoard()
	var gui GUI
	gui.Boards = g.Boards
	ctx.JSON(200, gui)
}

//LoadBoard .
func (g *Game) LoadBoard() {
	//Board Name
	if g.Director {
		name := g.SetBoardName("")
		g.WriteToPlayer(name)
		g.ListenToPlayer()
	} else {
		res, err := g.ListenToPlayer()
		if err != nil {
			log.Fatal(err)
		}
		name := g.SetBoardName(res[0])
		g.WriteToPlayer(name)
	}
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
	ABC = strings.Trim(ABC, nameTaken)
	name = append(name, nameTaken)
	for i := range g.Boards {
		g.Boards[i].Name = string([]rune(ABC)[i])
		name[0] = name[0] + string([]rune(ABC)[i])
	}
	return name
}

//Init .
func (g *Game) Init( /*ctx *gin.Context*/ ) {
	g.Protocol.Reset()
	if !g.Director {
		g.Wait()
	}
	//ctx.JSON(http.StatusOK, "OK")
}

//Update .
func (g *Game) Update(ctx *gin.Context) {
	var gui GUI
	if !g.Message.GetMessageFinished() {
		if g.Director {
			if g.Message.Bingo == "null" {
				ball := g.Blower.GetBallOut()
				g.Play(ball)
				g.Message.SaveBall(ball)
			} else {
				g.Message.Finished = "true"
			}
		} else if g.Message.Ball != "null" {
			g.Play(g.Message.GetMessageBall())
		}
	} else {
		gui.Bingo = g.Message.GetMessageBingo()
	}
	gui.Ball = g.Message.GetMessageBall()
	gui.Boards = g.Boards
	ctx.JSON(200, gui)
}

//Play .
func (g *Game) Play(ball items.Ball) {
	bingo := false
	for i := range g.Boards {
		row, column := g.Boards[i].Take(ball)
		if row != -1 && column != -1 {
			if g.Mode == "lineal" {
				bingo = g.Boards[i].CheckBoardLine(row, column)
			} else {
				bingo = g.Boards[i].CheckFull()
			}
			if bingo {
				g.Message.SaveWinner(g.Boards[i].Name)
			}
		}
	}
}

//Send .
func (g *Game) Send( /*ctx *gin.Context*/ ) {
	var message []string
	g.Protocol.Reset()
	message = append(message, g.Message.Ball)
	message = append(message, g.Message.Bingo)
	message = append(message, g.Message.Finished)
	g.WriteToPlayer(message)
	g.Protocol.Reset()
	//ctx.JSON(http.StatusOK, "OK")
}

//Wait .
func (g *Game) Wait( /*ctx *gin.Context*/ ) {
	res, err := g.ListenToPlayer()
	if err != nil {
		log.Fatal(err)
	}
	g.Message.SaveMessage(res)
	//ctx.JSON(http.StatusOK, "OK")
}

//Close .
func (g *Game) Close( /*ctx *gin.Context*/ ) {
	g.Protocol.Close()
	//ctx.JSON(http.StatusOK, "OK")
}
