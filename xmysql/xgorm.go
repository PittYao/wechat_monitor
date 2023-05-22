package xmysql

import (
	"fmt"
	"github.com/PittYao/wechat_monitor/etc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type XGorm struct {
	DB *gorm.DB
}

var G *XGorm

func NewXGorm() *XGorm {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		etc.C.Database.Username, etc.C.Database.Password, etc.C.Database.Host, etc.C.Database.Port, etc.C.Database.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: true,
			},
		),
	})
	if err != nil {
		fmt.Println("Gorm connect mysql err:", err)
		panic(err)
	}

	fmt.Println("Gorm Running")
	return &XGorm{DB: db}
}

func (gu *XGorm) Close() error {
	db, err := gu.DB.DB()
	if err != nil {
		fmt.Println("close Gorm err:", err)
		return err
	}

	return db.Close()
}
