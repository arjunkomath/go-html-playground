package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Title": "Todos!",
		})
	})

	router.GET("/todos", func(c *gin.Context) {
		todos := make([]string, 0)
		todos = append(todos, "Todo 1")
		todos = append(todos, "Todo 2")

		c.HTML(http.StatusOK, "todos.tmpl", gin.H{
			"Todos": todos,
		})
	})

	router.Run()
}
