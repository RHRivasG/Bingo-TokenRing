package main

import (
	"bingo-tokenring/logic"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/fatih/color"
)

//G .
var G logic.Game

func main() {
	//Parametros
	// 1 -> TCP
	// 2 -> Listener
	// 3 -> Writer
	// 4 -> Mode
	// 5 -> Cartones
	i, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatal(err)
	}
	G = logic.NewGame(os.Args[2], os.Args[3], i, os.Args[4])
	gui := G.LoadGame()
	RunGame(gui)
	G.Init()
	//var b []byte = make([]byte, 1)
	for len(gui.Bingo) == 0 {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		gui = G.Update()
		RunGame(gui)
		fmt.Println()
		fmt.Println("Cartones ganadores: ", gui.Bingo)
		//os.Stdin.Read(b)
		time.Sleep(1 * time.Second)
		//Read()
		G.Send()
		G.Wait()

	}
	defer G.Close()
}

//RunGame .
func RunGame(g logic.GUI) {
	var valor string
	fmt.Println("Pelota: ", g.Ball.Letter, g.Ball.Number)
	for _, board := range g.Boards {
		fmt.Println("Carton: ", board.Name)
		color.Yellow(" [B ]  [I ]  [N ]  [G ]  [O ]")
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				valor = strconv.Itoa(board.Tiles[i][j].Number)
				if board.Tiles[i][j].Number < 10 {
					valor = "0" + valor
				}
				if board.Tiles[i][j].Taken {
					c := color.New(color.FgCyan).Add(color.Underline)
					c.Print(" [", valor, "] ")
				} else {

					fmt.Print(" [", valor, "] ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
