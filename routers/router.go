package routers

import (
	"LetsGo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	// JS, CSS, Images -> stored in static file fold
	// HTML -> stored in templates
	engine.LoadHTMLGlob("templates/*.html")
	engine.GET("/", controller.ShowLists)
	// use POST Method: front end use a POST form
	// use GET Method: others
	listGroup := engine.Group("/list")
	{
		listGroup.POST("/create", controller.CreateList)
		listGroup.GET("/delete/:id", controller.DeleteList)
		listGroup.GET("/complete/:id", controller.CompleteList)
	}
	return engine
}
