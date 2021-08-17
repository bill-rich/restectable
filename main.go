package main

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

var messages map[string]string

type Message struct {
	Content string `json:"content"`
	Hash    string `json:"hash"`
}

func main() {
	apiServer := setupRouter()
	apiServer.Run()
}

func setupRouter() *gin.Engine {
	messages = map[string]string{}
	apiServer := gin.Default()
	apiServer.POST("/message", createMessage)
	apiServer.GET("/message/:hash", getMessage)
	return apiServer
}

func createMessage(c *gin.Context) {
	var message Message
	c.Bind(&message)
	hashHex := sha256.Sum256([]byte(message.Content))
	hash := hex.EncodeToString(hashHex[:])

	messages[hash] = message.Content
	c.JSON(200, gin.H{
		"hash": hash,
	})
}

func getMessage(c *gin.Context) {
	hash := c.Param("hash")
	c.JSON(200, gin.H{
		"content": messages[hash],
	})
}
