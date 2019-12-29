package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

)

type Bumbu struct {
	Ingredients []struct {
		Title      string `json:"title"`
		BestBefore string `json:"best-before"`
		UseBy      string `json:"use-by"`
	} `json:"ingredients"`
}
type Resep struct {
	Recipes []struct {
		Title       string   `json:"title"`
		Ingredients []string `json:"ingredients"`
	} `json:"recipes"`
}

type DataLunch struct {
	Title       string   `json:"title"`
	Ingredients []DataIngredient `json:"ingredients"`
}

var (
	dbRecipe     = initRecipe()
	dbIngredient = initIngredients()
)

func ShowRecipe(c *gin.Context) {
	var d Resep
	if err := json.Unmarshal([]byte(dbRecipe), &d); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, d)
}

func ShowIngredient(c *gin.Context) {
	var d Bumbu
	if err := json.Unmarshal([]byte(dbIngredient), &d); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, d)
}

type DataIngredient struct{
	Ingredients string `json:"ingredients"`
	BestBefore  string   `json:"best-before"`
	UseBy       string   `json:"use-by"`
}
func ShowLunch(c *gin.Context) {
	var data []DataIngredient
	var lunch []DataLunch
	
	menu := c.Params.ByName("menu")
	
	var d Resep
	if err := json.Unmarshal([]byte(dbRecipe), &d); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	jml := len(d.Recipes)
	for i := 0; i < jml; i++ {
		titles := d.Recipes[i].Title
		
		if titles == menu {

		var b Bumbu
		if err := json.Unmarshal([]byte(dbIngredient), &b); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		t := len(d.Recipes[i].Ingredients)
		for tt := 0; tt < t; tt++ {
			ing := fmt.Sprint(d.Recipes[i].Ingredients[tt])
			jing := len(b.Ingredients)
			
			for a := 0; a < jing; a++ {
				ingb := fmt.Sprint(b.Ingredients[a].Title)
				bestb := b.Ingredients[a].BestBefore
				useb := b.Ingredients[a].UseBy
				if ing == ingb {
					ct := DataIngredient{Ingredients: fmt.Sprint(d.Recipes[i].Ingredients[tt]), BestBefore: bestb, UseBy: useb}
					data = append(data, ct)
				}
			}
		}
		ctx := DataLunch{Title : titles,Ingredients : data}
		lunch = append(lunch,ctx)
		}
		
	}
	c.JSON(http.StatusOK, lunch)
}
