package manager

import (
	"book/config"
	"book/structs"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Save(user *structs.BookInfo)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	config.InitConfig()
	dsn := viper.GetString("datasource.DatabaseURI")
	fmt.Println(1111)
	fmt.Println(viper.GetString("datasource.DatabaseURI"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&structs.BookInfo{})
}

func (mgr manager) Save(BookInfo *structs.BookInfo) {
	mgr.db.Create(BookInfo)
}
