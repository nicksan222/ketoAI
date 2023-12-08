package preferences_follows_list

import "context"

type PreferencesFollowsListRequest struct {
	UserId  string `json:"userId"`
	Context context.Context
}
