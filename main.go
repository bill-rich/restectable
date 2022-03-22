package main

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Content string `json:"content"`
	Hash    string `json:"hash"`
}

func main() {
	apiServer := setupRouter()
	apiServer.Run()
}

func setupRouter() *gin.Engine {
	messages := map[string]string{}
	apiServer := gin.Default()
	apiServer.POST("/message", createMessageFunc(messages))
	apiServer.GET("/message/:hash", getMessageFunc(messages))
	return apiServer
}

func createMessageFunc(messages map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		var message Message
		c.Bind(&message)
		hashHex := sha256.Sum256([]byte(message.Content))
		hash := hex.EncodeToString(hashHex[:])

		messages[hash] = message.Content
		c.JSON(200, gin.H{
			"hash": hash,
		})
	}
}

func getMessageFunc(messages map[string]string) func(*gin.Context) {
	return func(c *gin.Context) {
		hash := c.Param("hash")
		if _, ok := messages[hash]; !ok {
			c.JSON(404, nil)
			return
		}
		c.JSON(200, gin.H{
			"content": messages[hash],
		})
	}
}
