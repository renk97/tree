package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/router"
)

func main() {
	r := gin.Default()
	r.GET("/get_tree", router.GetTreeRouter)
	r.Run(":80")
}
