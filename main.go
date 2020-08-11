package main

import (
	"bingo-tokenring/logic"
	"fmt"
	"os"
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
				protocol.Write(messages)
			}
			if messages[0] == protocol.GetWriterName() {
				fmt.Println("H")
				protocol.Write([]string{"Hola"})
			} else {
				res, err := protocol.Listen()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(res)
			}

		} else {
			fmt.Println("No hay argumentos")
		}
	*/
	/*
		i, err := strconv.Atoi(os.Args[5])
		if err != nil {
			log.Fatal(err)
		}
		g := logic.NewGame(os.Args[2], os.Args[3], i, os.Args[4])
		r := gin.Default()
		r.Use(cors.Default())

		//Routes
		r.GET("init", func(ctx *gin.Context) {
			g.Init()
		})
		r.GET("loadgame", func(ctx *gin.Context) {
			g.LoadGame()
		})
		r.GET("update", func(ctx *gin.Context) {
			g.Update(ctx)
		})
		r.GET("send", func(ctx *gin.Context) {
			g.Send()
		})
		r.GET("wait", func(ctx *gin.Context) {
			g.Wait()
		})

		r.Run(":" + os.Args[1])
	*/
	g := logic.NewGame(os.Args[2], os.Args[3], 2, "lineal")
	g.LoadGame()
	g.Close()
	fmt.Println(g.Boards)
}
