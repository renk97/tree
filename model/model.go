package model

import (
	"encoding/json"
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

func CreateTreeModel(input IOTree) (err error) {
	// 整個資料轉json帶入
	json_leaves, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return
	}

	data := Tree{
		Id:     input.Id,
		Root:   input.Root,
		Leaves: json_leaves,
	}

	err = db.Table("Mtree").Create(&data).Error

	return
}

func UpdateLeafModel(input IOTree) (err error) {
	json_leaves, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return
	}

	data := Tree{
		Root:   input.Root,
		Leaves: json_leaves,
	}

	err = db.Table("Mtree").Where("id = ?", input.Id).Updates(&data).Error

	return
}

func DeleteTreeModel(id int, root string) (err error) {
	raw := db.Table("Mtree")

	if id != 0 {
		raw = raw.Where("id = ?", id)
	} else if root != "" {
		raw = raw.Where("root = ?", root)
	}

	err = raw.Delete(&Tree{}).Error

	return
}

func CreateHashTreeModel(input HashIOTree) (err error) {
	//var raw *gorm.DB
	var data Tree
	// 整個資料轉json帶入
	json_leaves, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return
	}

	data = Tree{
		Id:     input.Id,
		Root:   input.Root,
		Leaves: json_leaves,
	}

	err = db.Table("Mtree_hash").Create(&data).Error

	return
}

func GetTreeModel(leaf_type string, id int, root string) (t []Tree, err error) {
	table := "Mtree"
	if leaf_type == "hash" {
		table = "Mtree_hash"
	}

	raw := db.Table(table)

	if id != 0 {
		raw = raw.Where("id = ?", id)
	} else if root != "" {
		raw = raw.Where("root = ?", root)
	}

	err = raw.Find(&t).Error

	return
}
