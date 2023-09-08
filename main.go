package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Tree struct {
	Id     int
	Root   string
	Struct string
}

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(dsn)

	if err != nil {
		log.Println(err)
	}

	db_pool, err := db.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return
	}

	db_pool.SetMaxIdleConns(10)
	db_pool.SetMaxOpenConns(100)
}

func main() {
	r := gin.Default()

	r.GET("/get_tree", GetTreeRouter)
	r.Run(":80")
}

func GetTreeRouter(c *gin.Context) {
	resp, code := GetTreeController()

	c.JSON(http.StatusOK, gin.H{
		"tree_data": resp,
		"status":    code,
	})
}

func GetTreeController() (Tree, int) {
	code := http.StatusOK

	resp, err := GetTreeModel()
	if err != nil {
		code = http.StatusForbidden
	}

	return resp, code
}

func GetTreeModel() (Tree, error) {
	t := Tree{}

	err := db.Table("test_tree").Find(&t).Error

	return t, err
}
