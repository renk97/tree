package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"MTree/controller"
	"MTree/model"
)

// 新增樹
func CreateTreeRouter(c *gin.Context) {
	var input model.IOTree
	c.Bind(&input)

	code := controller.CreateTreeController(input)

	c.JSON(code, gin.H{
		"tree_data": "",
		"status":    code,
	})
}

// 修改葉節點
func UpdateLeafRouter(c *gin.Context) {
	var input model.IOTree
	c.Bind(&input)

	code := controller.UpdateLeafController(input)

	c.JSON(code, gin.H{
		"tree_data": "",
		"status":    code,
	})
}

// 刪除樹
func DeleteTreeRouter(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	root := c.Query("root")

	code := controller.DeleteTreeController(id, root)

	c.JSON(code, gin.H{
		"tree_data": "",
		"status":    code,
	})
}

// 新增 hash leaf 樹
func CreateHashTreeRouter(c *gin.Context) {
	var input model.HashIOTree
	c.Bind(&input)

	code := controller.CreateHashTreeController(input)

	c.JSON(code, gin.H{
		"tree_data": "",
		"status":    code,
	})
}

// 取資料
func GetTreeRouter(c *gin.Context) {
	leaf_type := c.Query("lt")
	id, _ := strconv.Atoi(c.Query("id"))
	root := c.Query("root")

	resp, code := controller.GetTreeController(leaf_type, id, root)

	c.JSON(code, gin.H{
		"tree_data": resp,
		"status":    code,
	})
}
