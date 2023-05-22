package xgin

import (
	"fmt"
	"github.com/PittYao/wechat_monitor/etc"
	"github.com/PittYao/wechat_monitor/model"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	_ "time/tzdata"
)

func NewGin() {
	// 创建 Gin 引擎
	engine := gin.Default()

	// 添加全局中间件
	engine.Use(GlobalMiddleware())
	engine.Use(CORSMiddleware())

	// 创建自定义路由组
	api := engine.Group("/api")
	{
		api.GET("/list", getWechatHandler)
	}

	// 启动应用
	engine.Run(":" + etc.C.Port)
}

// 全局中间件示例
func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印请求入参
		fmt.Println("Request:", c.Request.URL.Path, c.Request.Method, c.Request.URL.Query())

		// 添加 defer recover 代码块
		defer func() {
			if r := recover(); r != nil {
				// 恢复处理逻辑，例如打印错误信息或返回错误响应
				fmt.Println("Recovered from panic:", r)
				c.JSON(500, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		// 调用 Next() 继续处理请求
		c.Next()

		// 执行一些后置处理逻辑
		// ...
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// 处理 GET /api/list 请求
func getWechatHandler(c *gin.Context) {
	var (
		queryStr  string
		pageStr   string
		startTime time.Time
		endTime   time.Time
	)

	// 查询参数
	queryStr = c.Query("q")

	// 分页参数
	pageStr = c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page"})
		return
	}

	// 获取时间参数
	yearStr := c.Query("year")
	if yearStr != "" {
		// 解析年份参数
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid year"})
			return
		}

		// 构建开始时间和结束时间
		startTime = time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
		endTime = time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.Local)
	} else {
		startTime = time.Time{}
	}

	// 查询数据
	weChats, err := model.GetData(queryStr, startTime, endTime, page)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 返回查询结果
	c.JSON(200, gin.H{"data": weChats})

}
