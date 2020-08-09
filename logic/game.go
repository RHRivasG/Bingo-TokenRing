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
}

//Message .
type Message struct {
	Blower string
	Ball   string
	Bingo  string
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
		g.LastContactPlayer(message)
	} else {
		g.LastContactPlayer(res)
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
		res, err = g.ListenPlayer()
		if err != nil {
			log.Fatal(err)
		}
		req := g.SetBoardName(res[0])
		g.LastContactPlayer(req)
	}

	//GUI
	var gui GUI
	gui.Boards = g.Boards

	return ctx.JSON(200, gui)
}

//ListenPlayer .
func (g *Game) ListenPlayer() ([]string, error) {
	return g.Protocol.Listen()
}

//ContactPlayer .
func (g *Game) ContactPlayer(message []string) ([]string, error) {
	res, err := g.Protocol.Converse(message)
	return res, err
}

//LastContactPlayer .
func (g *Game) LastContactPlayer(message []string) {
	g.Protocol.EndConversation(message)
}

//SetBoardName .
func (g *Game) SetBoardName(nameTaken string) []string {
	var name []string
	name = append(name, "")
	strings.Trim(ABC, nameTaken)
	for i, b := range g.Boards {
		b.Name = string([]rune(ABC)[i])
		name[0] = name[0] + string([]rune(ABC)[i])
	}
	return name
}

//Validate .
func (g *Game) Validate(messages []string) {
	//llenar bingo
	message := Message{messages[0], messages[1], messages[2]}
	if message.Blower != g.Protocol.GetWriterName() {
		if message.Ball != "bingo" {
			if message.Ball != "null" {
				var ball items.Ball
				//convertir string a ball

				row, column := g.Boards[0].Take(ball)
				if g.Boards[0].CheckBoardLine(row, column) {
				}

				//Chequear si hay bingo
				//Si hay bingo, concatenar
			}
		} else {
			//Termino ronda, decir ganadores
			//Hacer funcion que separe message[2]
			//pasarlo a GGUI.Bingo
		}

	} else {
		if message.Ball != "bingo" {
			if message.Ball != "null" {
				//Marcar tablero
				//Chequear si hay bingo
				//Si hay bingo, concatenar
			}
		} else {
			//Termino ronda, decir ganadores
			//Hacer funcion que separe message[2]
			//pasarlo a GGUI.Bingo
		}
		//if hay bingo, ball = null,
		//Sino sacar pelota
	}
	//pasar mensaje
}
