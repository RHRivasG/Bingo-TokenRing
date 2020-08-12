package main

import (
	"bingo-tokenring/logic"
	"bufio"
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
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	y := color.New(color.FgYellow).Add(color.Underline)
	c := color.New(color.FgCyan).Add(color.Underline)
	g := color.New(color.FgGreen).Add(color.Underline)
	y.Println(" ______   _____  ____  _____   ______     ___    ")
	y.Println("|_   _ \\ |_   _||_   \\|_   _|.' ___  |  .'   `.  ")
	y.Println("  | |_) |  | |    |   \\ | | / .'   \\_| /  .-.  \\ ")
	y.Println("  |  __'.  | |    | |\\ \\| | | |   ____ | |   | | ")
	y.Println(" _| |__) |_| |_  _| |_\\   |_\\ `.___]  |\\  `-'  / ")
	y.Println("|_______/|_____||_____|\\____|`._____.'  `.___.'  ")
	c.Println(" _______     ______    _____   ______    _____   ")
	c.Println("|_   __ \\  .' ____ \\  / ___ `./ ____ `. / ___ `. ")
	c.Println("  | |__) | | (___ \\_||_/___) |`'  __) ||_/___) | ")
	c.Println("  |  __ /   _.____`.  .'____.'_  |__ '. .'____.' ")
	c.Println(" _| |  \\ \\_| \\____) |/ /_____| \\____) |/ /_____  ")
	c.Println("|____| |___|\\______.'|_______|\\______.'|_______| ")
	c.Println("                                                 ")
	g.Print("Esperando jugadores...")
	G = logic.NewGame(os.Args[2], os.Args[3], i, os.Args[4])
	gui := G.LoadGame()
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	RunGame(gui)
	G.Init()
	reader := bufio.NewReader(os.Stdin)
	for len(gui.Bingo) == 0 {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		gui = G.Update()
		RunGame(gui)
		g.Print("Cartones ganadores: ", gui.Bingo)
		if os.Args[6] == "manual" {
			reader.ReadString('\n')
		} else {
			time.Sleep(2 * time.Second)
		}
		G.Send()
		G.Wait()

	}
	defer G.Close()
}

//RunGame .
func RunGame(g logic.GUI) {
	var valor string
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Pelota: ", g.Ball.Letter, g.Ball.Number)
	fmt.Println()
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
