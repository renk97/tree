package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/controller"
)

func GetTreeRouter(c *gin.Context) {
	resp, code := controller.GetTreeController()

	c.JSON(http.StatusOK, gin.H{
		"tree_data": resp,
		"status":    code,
	})
}
