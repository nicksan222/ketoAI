package recipes

// Required for handling MongoDB ObjectID

type CreateRecipeRequest struct {
	Steps       []string `json:"steps"`
	Tags        []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
}
