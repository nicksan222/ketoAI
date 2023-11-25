package ingredients_list

type ListIngredientsRequest struct {
	BeginsWith string `json:"begins_with"`
	EndsWith   string `json:"ends_with"`
	Limit      int64  `json:"limit"`
}
