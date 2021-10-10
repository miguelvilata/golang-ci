package main

import (
	"os"

	"github.com/rubencougil/geekshubs-library/pkg/api"
	library_pkg "github.com/rubencougil/geekshubs-library/pkg/library"

	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-gonic/gin"
)

func main() {

	logger := logrus.New()
	dbConn := os.Getenv("DB")
	library := library_pkg.NewLibrary(library_pkg.NewMysqlRepository(dbConn))
	handlers := api.NewAPI(logger)

	router := gin.New()
	router.Use(ginlogrus.Logger(logger.WithField("service", "gin")), gin.Recovery())
	api := router.Group("/api")

	logger.Info("Welcome to the GeeksHubs Library. API is ready.")

	api.GET("/", handlers.Hello)
	api.POST("/books", handlers.AddBook(library))

	api.GET("/books", handlers.GetAllBooks(library))
	api.GET("/books/:id", handlers.GetBook(library))
	api.DELETE("/books/:id", handlers.DeleteBook(library))
	api.PUT("/books/:id", handlers.UpdateBook(library))

	_ = router.Run(":8080")
}
