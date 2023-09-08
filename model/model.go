package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

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

func GetTreeModel(id int, root string) (t []Tree, err error) {
	raw := db.Table("Mtree")

	if id != 0 {
		raw = raw.Where("id = ?", id)
	}
	if root != "" {
		raw = raw.Where("root = ?", root)
	}

	err = raw.Find(&t).Error

	return
}

// func CreateTreeModel() (t Tree, err error) {
// 	err = db.Table("Mtree").Find(&t).Error

// 	return
// }
