package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	r := gin.Default()

	r.GET("/random/string", func(c *gin.Context) {
		randomString := RandomString(10)
		c.JSON(http.StatusOK, gin.H{"quote": randomString})
	})

	r.GET("/random/integer", func(c *gin.Context) {
		randomNumber := rand.Intn(1001)
		c.JSON(http.StatusOK, gin.H{"quote": randomNumber})
	})

	err := r.Run("0.0.0.0:5000")
	if err != nil {
		panic(err)
	}
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
