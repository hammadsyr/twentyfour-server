package main

import (
	"github.com/gin-gonic/gin"
	"twentyfour.com/server/play"
)

func main() {
	play.ConnectMongodb()
	router := gin.Default()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	router.GET("/join", play.Join)
	router.GET("/leaderboard", play.GetLeaderboard)
	router.POST("/create")
	router.Run()
}
