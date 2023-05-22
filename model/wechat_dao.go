package model

import (
	"errors"
	"github.com/PittYao/wechat_monitor/xmysql"
	"gorm.io/gorm"
	"time"
)

var (
	pageSize = 20
)

func DataHandler(d Data) {
	// 根据aid查询数据是否存在
	for _, article := range d.Articles {
		var w Wechat
		xmysql.G.DB.Where("aid = ?", article.Aid).First(&w)

		if w.ID == 0 {
			// 数据不存在 处理数据
			// CreateTime 由时间戳转为 字符串格式
			if article.CreateTime != 0 {
				article.CreateTimeTr = TimestampConvert(article.CreateTime)
				article.CreateTimeStr = TimeFormat(article.CreateTimeTr)
				xmysql.G.DB.Save(article)
			}
		}
	}
}

func TimestampConvert(timestamp int64) time.Time {
	// 使用 time.Unix() 函数将时间戳转换为时间对象
	t := time.Unix(timestamp, 0)
	return t
}
func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetData(queryStr string, startTime, endTime time.Time, page int) ([]Wechat, error) {
	var (
		weChats []Wechat
		tx      *gorm.DB
	)

	if page == 0 {
		page = 1
	}

	tx = xmysql.G.DB.Order("create_time_tr desc")

	// 关键词查询
	if queryStr != "" {
		tx = tx.Where("title like ?", "%"+queryStr+"%")
	}

	// 年份查询
	if !startTime.IsZero() && !endTime.IsZero() {
		tx = tx.Where("create_time_tr >= ? AND create_time_tr < ?", startTime, endTime)
	}

	// 分页查询
	tx.Limit(pageSize).Offset((page - 1) * pageSize)

	tx.Find(&weChats)
	if tx.Error != nil {
		return nil, errors.New("数据查询异常")
	}
	return weChats, nil
}
