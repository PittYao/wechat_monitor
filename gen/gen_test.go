package gen

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testing"
)

func TestGenStruct(t *testing.T) {
	dsn := "root:root@(127.0.0.1:3306)/wechat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./model",
	})
	g.UseDB(db)
	// 填要生成的表名
	g.GenerateModel("wechat")
	g.Execute()
}
