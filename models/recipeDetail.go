package models

type MeasureType int

const (
	_ MeasureType = iota
	Cup
	Tablespoon
	Teaspoon
	Gram
	Mg
	Litre
	Ml
	Ounce
)

type RecipeDetail struct {
	Id               int         `gorm:"primaryKey" json:"id"`
	RecipeId         int         `json:"recipeId"`
	IngredientId     int         `json:"ingredientId"`
	IngredientStatus string      `json:"ingredientStatus"`
	MeasureType      MeasureType `json:"measureType"`
	MeasureQty       float32     `json:"measureQty"`
	Recipe           *Recipe     `gorm:"ForeignKey:Id;AssociationForeignKey:RecipeId" json:"recipe,omitempty"`
	Ingredient       *Ingredient `gorm:"ForeignKey:Id;AssociationForeignKey:IngredientId" json:"ingredient,omitempty" `
}
