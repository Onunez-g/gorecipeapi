package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/onunez-g/gorecipeapi/data"
	"github.com/onunez-g/gorecipeapi/models"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var recipes []models.Recipe

	var db *gorm.DB = data.Db

	for k, v := range params {
		if len(v) != 0 {
			queryString := k + "= ?"
			db = db.Where(queryString, strings.Join(v, " "))
		}
	}

	err := db.Preload("Details").Preload("Details.Ingredient").Find(&recipes).Error
	if err != nil {
		log.Println(err.Error())
	}

	response, err := json.Marshal(&recipes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe

	data.Db.Preload("Details").Preload("Details.Ingredient").First(&recipe, params["id"])

	if recipe.Id == 0 {
		w.WriteHeader(http.StatusPreconditionFailed)
		log.Println("Recipe not found")
		w.Write([]byte("Recipe not found"))
		return
	}

	response, err := json.Marshal(&recipe)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
	}

	data.Db.Create(&recipe)

	response, err := json.Marshal(&recipe)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe
	var updatedRecipe models.Recipe

	err := json.NewDecoder(r.Body).Decode(&updatedRecipe)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
	}

	data.Db.First(&recipe, params["id"]).Updates(&updatedRecipe)

	response, err := json.Marshal(&recipe)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipe models.Recipe
	var recipes []models.Recipe

	data.Db.First(&recipe, params["id"])
	data.Db.Delete(&recipe)
	data.Db.Find(&recipes)

	response, err := json.Marshal(&recipes)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateRecipeDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipeDetail models.RecipeDetail
	var updatedRecipeDetail models.RecipeDetail

	err := json.NewDecoder(r.Body).Decode(&updatedRecipeDetail)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
	}

	data.Db.First(&recipeDetail, params["id"]).Updates(&updatedRecipeDetail)

	response, err := json.Marshal(&recipeDetail)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DeleteRecipeDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var recipeDetail models.RecipeDetail
	var details []models.RecipeDetail

	data.Db.First(&recipeDetail, params["id"])
	data.Db.Delete(&recipeDetail)
	data.Db.Where("recipe_id = ?", recipeDetail.RecipeId).Find(&details)

	response, err := json.Marshal(&details)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
