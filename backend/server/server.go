package server

import (
	"log"

	"github.com/erwan690/blog/backend/config"
	"github.com/erwan690/blog/backend/db"
	_ "github.com/erwan690/blog/backend/docs"
	"github.com/erwan690/blog/backend/handler"
	"github.com/erwan690/blog/backend/repo"
	"github.com/erwan690/blog/backend/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	config := config.LoadConfig()
	dbc, err := db.GetDatabaseConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// init repo and service
	repo := repo.NewPostRepo(dbc)
	service := service.NewPostService(repo)
	// init handler
	handler := handler.NewPostHandler(service)

	r := NewRouter()

	article := r.Group("/article")

	article.GET("/", handler.GetAllPosts)
	article.GET("/:id", handler.GetPostByID)
	article.POST("/", handler.CreatePost)
	article.PUT("/:id", handler.UpdatePost)
	article.DELETE("/:id", handler.DeletePost)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":" + config.APPPORT)
}
