package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/junior/products-api/cmd/server/handler"
	"github.com/junior/products-api/internal/products"
	"github.com/junior/products-api/pkg/store"

	"github.com/junior/products-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TokenMiddleware(ctx *gin.Context) {
	tokenEnvironment := os.Getenv("TOKEN")
	token := ctx.GetHeader("token")
	if token != tokenEnvironment {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token inválido",
		})
		return
	}
	ctx.Next()
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsofservice https://developers.mercadolibre.com.ar/es_ar/terminos-y-condicianes
// @contact.name API Support
// @contact.url https: //developers.mercadolibre. com.ar/support
// @license.nane Apache 2.0
// @icense.url http: //wiw.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro ao carregar arquivo .env, ", err)
	}

	db := store.NewFileStore("file", "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	productHandler := handler.NewProduct(service)

	server := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Use(TokenMiddleware)

	pr := server.Group("/products")
	pr.POST("/", productHandler.Create())
	pr.GET("/", productHandler.GetAll())
	pr.GET("/:id", productHandler.GetByID())
	pr.PUT("/:id", productHandler.Update())
	pr.PATCH("/:id", productHandler.UpdateFields())
	pr.DELETE("/:id", productHandler.Delete())
	if err := server.Run(":4000"); err != nil {
		log.Fatal(err)
	}
}
