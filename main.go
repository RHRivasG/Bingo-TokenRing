package main

import (
	"bingo-tokenring/logic"
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	/*
		if os.Args[1:] != nil {
			protocol, err := protocol.NewProtocol(os.Args[1], os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			var arbitro []string
			arbitro = append(arbitro, os.Args[2])
			messages, err := protocol.Converse(arbitro)
			if err != nil {
				log.Fatal(err)
			}
			if messages[0] == protocol.GetWriterName() {
				fmt.Println(protocol)
			} else {
				fmt.Println(messages[0] + " es el arbitro")
				protocol.EndConversation(messages)
			}

		} else {
			fmt.Println("No hay argumentos")
		}
	*/
	i, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatal(err)
	}
	g := logic.NewGame(os.Args[2], os.Args[3], i, os.Args[4])
	e := echo.New()

	//Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//Routes
	e.GET("/loadgame", func(ctx echo.Context) error {
		return g.LoadGame(ctx)
	})

	e.Logger.Fatal(e.Start(os.Args[1]))
}
