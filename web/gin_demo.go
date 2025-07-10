package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.Query("name")
		c.DefaultQuery("age", "20")
		c.String(http.StatusOK, "hello World!")
	})
	r.GET("/news", func(c *gin.Context) {
		c.String(http.StatusOK, "这是GET新页面")
	})
	r.POST("/add", func(c *gin.Context) {
		c.PostForm("username")
		c.DefaultPostForm("age", "20")

		user := &User{}
		c.ShouldBind(&user)
		c.String(http.StatusOK, "这是POST请求路径")
	})
	r.PUT("/update", func(c *gin.Context) {
		c.String(http.StatusOK, "这是PUT请求路径")
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "这是DELETE请求路径")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"name": "张三",
			"age":  "18",
		})
	})

	r.GET("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "张三",
			"age":  "18",
		})
	})

	r.GET("/json3", func(c *gin.Context) {
		a := &User{Username: "张三", Password: "123456"}
		c.JSON(http.StatusOK, a)
	})

	// http://127.0.0.1:8000/jsonp?callback=abc
	// 相应 abc({})
	r.GET("/jsonp", func(c *gin.Context) {
		a := &User{Username: "张三", Password: "123456"}
		c.JSONP(http.StatusOK, a)
	})

	r.GET("/xml", func(c *gin.Context) {
		a := &User{Username: "张三", Password: "123456"}
		c.XML(http.StatusOK, a)
	})

	r.POST("/xml", func(c *gin.Context) {
		user := &User{}
		xmlSliceData, _ := c.GetRawData()
		xml.Unmarshal(xmlSliceData, &user)
		fmt.Println(string(xmlSliceData))
		a := &User{Username: "张三", Password: "123456"}
		c.XML(http.StatusOK, a)
	})

	r.GET("/cookie", func(c *gin.Context) {
		c.SetCookie("username", "张三", 3600, "/", "127.0.0.1", false, true)
		a := &User{Username: "张三", Password: "123456"}
		c.XML(http.StatusOK, a)
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
