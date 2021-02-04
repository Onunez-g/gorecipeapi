package models

type MealType int
type Difficulty int

const (
	Breakfast MealType = iota
	Lunch
	Supper
	Snack
)

const (
	Beginner Difficulty = iota
	Intermediate
	Advanced
)

type Recipe struct {
	Id               int            `gorm:"primaryKey" json:"id"`
	Title            string         `json:"title"`
	MealType         MealType       `json:"mealType"`
	ServesNum        int            `json:"servesNum"`
	Difficulty       Difficulty     `json:"difficulty"`
	PreparationSteps string         `json:"preparationSteps"`
	Details          []RecipeDetail `gorm:"ForeignKey:RecipeId;AssociationForeignKey:Id" json:"details"`
}
