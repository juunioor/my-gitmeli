package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/junior/products-api/cmd/server/handler"
	"github.com/junior/products-api/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()
	pr := server.Group("/products")
	pr.POST("/", productHandler.Store())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.GetByID())
	if err := server.Run(":4000"); err != nil {
		log.Fatal(err)
	}
}
