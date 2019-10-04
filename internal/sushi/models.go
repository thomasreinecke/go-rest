package sushi

// Roll is model for sushi
type Roll struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}
