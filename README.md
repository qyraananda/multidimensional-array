# Multidimensional-Array
Menampilkan data array multidimensi dengan menggunakan Go Lang. Sumber data berasal dari data JSON

# Framework
Menggunakan framework <a href="https://gin-gonic.com/">Gin-Gonic</a>

# Server
aktif server http://localhost:8071/api/v5

di main.go
<pre>
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
</pre>

Method yang digunakan adalah <strong>GET</strong>

v5.GET("/showrecipe", app.ShowRecipe) = menampilkan daftar resep
http://localhost:8071/api/v5/showrecipe

v5.GET("/showingredient", app.ShowIngredient) = menampilkan daftar bahan
http://localhost:8071/api/v5/showingredient

v5.GET("/showlunch/:menu",app.ShowLunch) = menampilkan daftar bahan untuk membuat resep
http://localhost:8071/api/v5/showlunch/Ham and Cheese Toastie

# Screenshot
<img src="https://github.com/qyraananda/multidimensional-array/blob/master/screenshot.png">
