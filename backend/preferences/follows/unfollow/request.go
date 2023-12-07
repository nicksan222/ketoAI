package follows_unfollow

type UnfollowRequest struct {
	UserId string `json:"user_id"`
	ToUnfollow string `json:"to_unfollow"`
}