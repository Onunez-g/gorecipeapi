package router

import (
	"github.com/gorilla/mux"
	"github.com/onunez-g/gorecipeapi/controllers"
)

func Get() *mux.Router {
	r := mux.NewRouter()

	//Ingredients Controller
	r.HandleFunc("/api/ingredients", controllers.GetIngredients).Methods("GET")
	r.HandleFunc("/api/ingredients/{id}", controllers.GetIngredient).Methods("GET")
	r.HandleFunc("/api/ingredients", controllers.CreateIngredient).Methods("POST")
	r.HandleFunc("/api/ingredients/{id}", controllers.UpdateIngredient).Methods("PUT")

	//Recipes Controller
	r.HandleFunc("/api/recipes", controllers.GetRecipes).Methods("GET")
	r.HandleFunc("/api/recipes/{id}", controllers.GetRecipe).Methods("GET")
	r.HandleFunc("/api/recipes", controllers.CreateRecipe).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", controllers.UpdateRecipe).Methods("PUT")
	r.HandleFunc("/api/recipes/{id}", controllers.DeleteRecipe).Methods("DELETE")
	r.HandleFunc("/api/recipes/detail/{id}", controllers.UpdateRecipeDetail).Methods("PUT")
	r.HandleFunc("/api/recipes/detail/{id}", controllers.DeleteRecipeDetail).Methods("DELETE")

	return r
}
