package follow_add_follow

type NewFollowRequest struct {
	UserId     string `json:"user_id" binding:"required"`
	ToFollow  string `json:"to_follow" binding:"required"`
}
