package app

import (
	"fmt"
	"io/ioutil"
	// "net/http"
// "path/filepath"
// "os"
)

func initRecipe()string {
	content, err := ioutil.ReadFile("./src/recipe/data.json")
	if err != nil {
		fmt.Println(err)
	}

	return string(content)
}

func initIngredients() string {
	content, err := ioutil.ReadFile("./src/ingredient/data.json")
	if err != nil {
		fmt.Println(err)
	}

	return string(content)
}