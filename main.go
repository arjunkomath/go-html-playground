package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	todos := make([]string, 0)
	todos = append(todos, "Some todo")

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
		c.HTML(http.StatusOK, "todos.tmpl", gin.H{
			"Todos": todos,
		})
	})

	router.POST("/todos", func(c *gin.Context) {
		todos = append(todos, c.PostForm("task"))

		c.HTML(http.StatusOK, "todos.tmpl", gin.H{
			"Todos": todos,
		})
	})

	router.Run()
}
