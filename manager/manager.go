package manager

import (
	"book/config"
	"book/structs"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type Manager interface {
	Save(user *structs.BookInfo)
	CreateZZ(name string)
	SaveZZ(name string, Chapter *structs.Chapter)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	config.InitConfig()
	dsn := viper.GetString("datasource.DatabaseURI")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&structs.BookInfo{})
}

func (mgr manager) Save(BookInfo *structs.BookInfo) {
	var u structs.BookInfo
	mgr.db.Where("title=?", BookInfo.Title).First(&u)
	fmt.Println(u.ID)
	if u.ID == 0 {
		mgr.db.Create(BookInfo)
	}

}

// 创建小说表
func (mgr manager) CreateZZ(name string) {
	var u structs.BookInfo
	mgr.db.Where("title=?", name).First(&u)
	s2 := strconv.Itoa(int(u.ID))
	//fmt.Println(name,"book_"+s2)
	//isFlag := mgr.db.Migrator().HasTable(s2)
	//fmt.Println(isFlag)
	mgr.db.Table("book_" + s2).AutoMigrate(&structs.Chapter{})
	//if !isFlag {
	//	mgr.db.Table(s2).AutoMigrate(&structs.Chapter{})
	//}
}

// 保存小说章节
func (mgr manager) SaveZZ(name string, Chapter *structs.Chapter) {
	var u structs.BookInfo
	mgr.db.Where("title=?", name).First(&u)
	s2 := strconv.Itoa(int(u.ID))
	fmt.Println(name, s2)
	mgr.db.Table("book_" + s2).Create(Chapter)
}
