package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vezzalinistefano/learning-htmx/models"
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
		query := ctx.Query("q")
		ctx.HTML(
			http.StatusOK,
			"index",
			gin.H{
				"title":   "Contacts",
				"payload": contactRepository.GetAll(query),
				"q":       query,
			},
		)
	})

	router.GET("/contacts/new", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"new_contact",
			gin.H{},
		)
	})

	router.POST("/contacts/new", func(ctx *gin.Context) {
		contact := &models.Contact{}
		if err := ctx.ShouldBind(contact); err != nil {
			ctx.String(http.StatusBadRequest, "Bad request: %v", err)
			return
		}
		contactRepository.InsertContact(*contact)
		ctx.Redirect(http.StatusFound, "/contacts")
	})

	router.Run(":8080")
}
