package main

import (
	"./app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

)

func SetupRouter() * gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router:= gin.Default()
	router.MaxMultipartMemory = 8 << 20
	v5:= router.Group("api/v5")
	
	{
		v5.GET("/showrecipe", app.ShowRecipe)
		v5.GET("/showingredient", app.ShowIngredient)
		v5.GET("/showlunch/:menu",app.ShowLunch)
	}
	return router
}
func main() {
	router:= SetupRouter()
	router.Use(cors.Default())

	router.Run(":8071")
}
