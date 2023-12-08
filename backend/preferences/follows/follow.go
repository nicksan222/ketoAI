package follows

type Follow struct {
	UserId        string   `json:"userId" bson:"user_id"`
	FollowedUsers []string `json:"followedUsers" bson:"followed_users"`
}
