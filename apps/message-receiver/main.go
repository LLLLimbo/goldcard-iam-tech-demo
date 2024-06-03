package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/rcv", func(c *gin.Context) {
		//print request body
		body, _ := c.GetRawData()
		log.Printf("Received message: %s  \n", body)
		c.JSON(http.StatusOK, gin.H{})
	})

	err := r.Run("0.0.0.0:17011")
	if err != nil {
		panic(err)
	}
}
