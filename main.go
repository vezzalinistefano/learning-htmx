package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/vezzalinistefano/learning-htmx/models"
	"github.com/vezzalinistefano/learning-htmx/repositories"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"add": func(a int, b int) int {
			return a + b
		},
	})
	router.LoadHTMLGlob("templates/*")

	contactRepository := repositories.ContactsRepository

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/contacts")
	})

	router.GET("/contacts", func(ctx *gin.Context) {
		query := ctx.Query("q")
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
		ctx.HTML(
			http.StatusOK,
			"index",
			gin.H{
				"title":   "Contacts",
				"payload": contactRepository.GetAll(query, page),
				"q":       query,
				"page":    page,
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

	router.GET("/contacts/:contact_id/view", func(ctx *gin.Context) {
		if contactID, err := strconv.Atoi(ctx.Param("contact_id")); err == nil {
			if contact, err := contactRepository.GetByContactID(contactID); err == nil {
				ctx.HTML(
					http.StatusOK,
					"view_contact",
					gin.H{
						"payload": contact,
					},
				)
			} else {
				ctx.AbortWithError(http.StatusNotFound, err)
			}
		} else {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

	})

	router.POST("/contacts/:contact_id/edit", func(ctx *gin.Context) {
		contact := &models.Contact{}
		if err := ctx.ShouldBind(contact); err != nil {
			ctx.String(http.StatusBadRequest, "Bad request: %v", err)
			return
		}

		if contactID, err := strconv.Atoi(ctx.Param("contact_id")); err == nil {
			contact.Id = contactID
			contactRepository.EditContact(*contact)
		}
		ctx.Redirect(http.StatusFound, "/contacts")
	})

	router.GET("/contacts/:contact_id/edit", func(ctx *gin.Context) {
		if contactID, err := strconv.Atoi(ctx.Param("contact_id")); err == nil {
			if contact, err := contactRepository.GetByContactID(contactID); err == nil {
				ctx.HTML(
					http.StatusFound,
					"edit_contact",
					gin.H{
						"payload": contact,
					},
				)
			} else {
				ctx.AbortWithError(http.StatusNotFound, err)
			}
		} else {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
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

	router.DELETE("/contacts/:contact_id/delete", func(ctx *gin.Context) {
		if contactID, err := strconv.Atoi(ctx.Param("contact_id")); err == nil {
			contactRepository.DeleteContactById(contactID)
		} else {
			ctx.AbortWithStatus(http.StatusNotFound)
		}

		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
