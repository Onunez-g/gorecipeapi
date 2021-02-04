package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onunez-g/gorecipeapi/data"
	"github.com/onunez-g/gorecipeapi/models"
)

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ingredients []models.Ingredient
	data.Db.Find(&ingredients)
	response, err := json.Marshal(&ingredients)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}
	w.Write(response)
}

func GetIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ingredient models.Ingredient
	data.Db.First(&ingredient, params["id"])

	if ingredient.Id == 0 {
		w.WriteHeader(http.StatusPreconditionFailed)
		log.Println("Ingredient not found")
		w.Write([]byte("Ingredient not found"))
		return
	}
	response, err := json.Marshal(&ingredient)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	var ingredient models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
	}

	data.Db.Create(&ingredient)

	response, err := json.Marshal(&ingredient)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ingredient models.Ingredient
	var updatedIngredient models.Ingredient

	err := json.NewDecoder(r.Body).Decode(&updatedIngredient)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
	}

	data.Db.First(&ingredient, params["id"]).Updates(&updatedIngredient)

	response, err := json.Marshal(&ingredient)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
