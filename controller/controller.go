package controller

import (
	"LetsGo/models"
	"github.com/gin-gonic/gin"
	"html"
	"net/http"
	"strconv"
)

func ShowLists(c *gin.Context) {
	lists, err := models.ShowLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to fetch database",
			"error":   err.Error()})
		return
	}
	// return the HTML page with json data stored in "lists"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"lists": lists,
	})
}

func CreateList(c *gin.Context) {
	// if front end use a Post From
	note := c.PostForm("item")
	// Escape special HTML characters in the string
	// Prevent HTML injection vulnerabilities (XSS cross-site scripting attacks) when outputting HTML
	list := models.List{
		Note: html.EscapeString(note),
	}
	if err := models.CreateList(&list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failure",
			"error":   err.Error(),
		})
	}
	// refresh the page
	ShowLists(c)
}

func DeleteList(c *gin.Context) {
	// get dynamic parameter
	idStr := c.Param("id")
	// convert string to int
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteList(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to delete a note",
			"error":   err.Error(),
		})
	}
	// refresh the page
	ShowLists(c)
}

func CompleteList(c *gin.Context) {
	// get dynamic parameter
	idStr := c.Param("id")
	// convert string to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to make a note completed",
			"error":   err.Error(),
		})
	}
	// IsComplete false -> true
	models.UpdateList(uint(id))
	// refresh the page
	ShowLists(c)
}
