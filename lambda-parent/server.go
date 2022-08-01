package main

import (
	"lambda-parent/api"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	r := gin.Default()

	corsEnv := os.Getenv("CORS")
	if corsEnv != "" {
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{corsEnv}
		r.Use(cors.New(config))
	}

	r.GET("/api/v1/posts/users", api.GetPostsAndUsers)
	r.GET("/api/v1/posts", api.GetPostsAndUsersTest)

	return r
}
