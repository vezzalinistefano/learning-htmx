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

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/contacts")
	})

	router.GET("/contacts", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"index",
			gin.H{
				"title":   "Contacts",
				"payload": contactRepository.GetAll(),
			},
		)
	})
	router.Run(":8080")
}
