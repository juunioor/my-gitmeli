package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		nome := "Junior"
		c.JSON(200, gin.H{
			"mensagem": "Olá, " + nome,
		})
	})

	router.Run()
}
