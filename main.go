package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/router"
)

func main() {
	r := gin.Default()
	r.GET("/get_tree", router.GetTreeRouter)
	r.POST("/create_tree", router.CreateTreeRouter)
	r.PATCH("/update_leaf", router.UpdateLeafRouter)
	r.Run(":80")
}
