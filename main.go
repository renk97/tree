package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Tree struct {
	Id     int
	Root   string
	Struct string
}

type OTree struct {
	Id     int
	Root   string
	Struct []string
}

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"),
		os.Getenv("DBNAME"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	fmt.Println(dsn)
	fmt.Printf("%+v\n", db)

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

func GetTreeController() (OTree, int) {
	code := http.StatusOK

	resp, err := GetTreeModel()
	struct_arr := strings.Split(resp.Struct, ",")
	output := OTree{
		Id:     resp.Id,
		Root:   resp.Root,
		Struct: struct_arr,
	}

	if err != nil {
		code = http.StatusInternalServerError
	}

	return output, code
}

func GetTreeModel() (t Tree, err error) {
	err = db.Table("Mtree").Find(&t).Error

	return
}
