package follow_add_follow

import "context"

type NewFollowRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	ToFollow string `json:"to_follow" binding:"required"`
	Context  context.Context
}
