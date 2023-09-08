package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/controller"
	"MTree/model"
)

func GetTreeRouter(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	root := c.Query("root")

	resp, code := controller.GetTreeController(id, root)

	c.JSON(code, gin.H{
		"tree_data": resp,
		"status":    code,
	})
}

func CreateTreeRouter(c *gin.Context) {
	var input model.IOTree
	c.Bind(&input)

	code := controller.CreateTreeController(input)

	c.JSON(code, gin.H{
		"tree_data": "",
		"status":    code,
	})
}
