package main

import (
	"bingo-tokenring/logic"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

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
		g.LoadGame(ctx)
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
	r.GET("close", func(ctx *gin.Context) {
		g.Close()
	})

	r.Run(":" + os.Args[1])
	/*
		g := logic.NewGame(os.Args[2], os.Args[3], 2, "lineal")
		g.Boards[0].Tiles[0][3].Taken = true
		g.Boards[0].Tiles[1][3].Taken = true
		g.Boards[0].Tiles[2][3].Taken = true
		g.Boards[0].Tiles[3][3].Taken = true
		g.Boards[0].Tiles[4][3].Taken = true
		g.Boards[1].Tiles[0][3].Taken = true
		g.Boards[1].Tiles[1][3].Taken = true
		g.Boards[1].Tiles[2][3].Taken = true
		g.Boards[1].Tiles[3][3].Taken = true
		g.Boards[1].Tiles[4][3].Taken = true
		g.LoadGame()
		g.Init()
		if g.Director {
			g.Update()
			fmt.Println(g.Message)
			g.Send()
			g.Wait()
			g.Update()
			fmt.Println(g.Message)
			g.Send()

		} else {
			fmt.Println(g.Message)
			g.Update()
			fmt.Println(g.Message)
			g.Send()
			g.Wait()
			g.Update()
			fmt.Println(g.Message)
			g.Send()
		}
		g.Close()*/
}
