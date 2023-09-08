package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/controller"
)

func GetTreeRouter(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	root := c.Query("root")

	resp, code := controller.GetTreeController(id, root)

	c.JSON(http.StatusOK, gin.H{
		"tree_data": resp,
		"status":    code,
	})
}

// func CreateTreeRouter(c *gin.Context) {
// 	input := model.IOTree
// 	c.Bind(&input)

// 	resp, code := controller.CreateTreeController(input)

// 	c.JSON(http.StatusOK, gin.H{
// 		"tree_data": resp,
// 		"status":    code,
// 	})
// }
