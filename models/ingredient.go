package models

type Ingredient struct {
	Id            int            `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	RecipeDetails []RecipeDetail `gorm:"ForeignKey:IngredientId;AssociationForeignKey:Id" json:"recipeDetails,omitempty"`
}
