package server

import (
	"go_gin_react_blog/internal/conf"
	"go_gin_react_blog/internal/store"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setRouter(cfg conf.Config) *gin.Engine {
	router := gin.Default()

	router.RedirectTrailingSlash = true

	if cfg.Env == "prod" {
		router.Use(static.Serve("/", static.LocalFile("./assets/build", true)))
	}

	api := router.Group("/api")
	api.Use(customErrors)
	{
		api.POST("/signup", gin.Bind(store.User{}), signUp)
		api.POST("/signin", gin.Bind(store.User{}), signIn)
	}

	authorized := api.Group("/")
	authorized.Use(authorization)
	{
		authorized.GET("/posts", indexPosts)
		authorized.POST("/posts", gin.Bind(store.Post{}), createPost)
		authorized.PUT("/posts", gin.Bind(store.Post{}), updatePost)
		authorized.DELETE("/posts/:id", deletePost)
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
