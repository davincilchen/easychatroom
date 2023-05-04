package main

// import (
// 	"log"
// 	"easychatroom/pkg/app/server"
// 	"easychatroom/pkg/config"
// )

// const confPath = "./config.json"

// func main() {

// 	cfg, err := config.New(confPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	svr := server.New(cfg)
// 	svr.Serve()

// }

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

var messages []Message

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// 獲取所有消息记录
	r.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, messages)
	})

	// 發送消息
	r.POST("/message", func(c *gin.Context) {
		var message Message
		err := c.BindJSON(&message)
		if err != nil {
			log.Println("Error binding message JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message JSON"})
			return
		}

		messages = append(messages, message)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 啟動服务器
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
