package recipes

// Recipe Represents a recipe
type Recipe struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

// Ingredient Represents individual ingredients
type Ingredient struct {
	Name string `json:"name"`
}
