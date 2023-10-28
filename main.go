package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

    "github.com/vezzalinistefano/learning-htmx/repositories"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

    contactRepository := repositories.ContactsRepository

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello World",
		})
	})

	router.GET("/contacts", func(c *gin.Context) {
        c.JSON(http.StatusOK, contactRepository.GetAll())
	})

	router.Run(":8080")
}
