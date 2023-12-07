package follow_add_follow_test

import (
	"context"
	"testing"

	follow_add_follow "github.com/nicksan222/ketoai/preferences/follows/follow"
	"github.com/stretchr/testify/assert"
)

type NewFollowTest struct {
	UserId     string
	ToFollow   string
	resultCode int
	res        follow_add_follow.NewFollowResponse
}

var NewFollowTests = []NewFollowTest{
	{
		UserId:     "test_add_follow_user_route",
		ToFollow:   "test_add_follow_user_route",
		resultCode: 200,
		res: follow_add_follow.NewFollowResponse{
			FollowId: "test_add_follow_user_route",
		},
	},
	{
		UserId:     "test_add_follow_user_route",
		ToFollow:   "test_add_follow_user_route_2",
		resultCode: 200,
		res: follow_add_follow.NewFollowResponse{
			FollowId: "test_add_follow_user_route_2",
		},
	},
	{
		UserId:     "test_add_follow_user_route",
		ToFollow:   "test_add_follow_user_route_3",
		resultCode: 200,
		res: follow_add_follow.NewFollowResponse{
			FollowId: "test_add_follow_user_route_3",
		},
	},
	{
		UserId:     "test_add_follow_user_route",
		ToFollow:   "test_add_follow_user_route",
		resultCode: 404,
		res:        follow_add_follow.NewFollowResponse{},
	},
}

func TestNewFollowHandler(t *testing.T) {
	for _, test := range NewFollowTests {
		id := createMockUserWithFollowers(t, test.UserId, []string{})
		defer deleteMockUserWithFollowers(t, id)

		// Test the handler
		res, err := follow_add_follow.NewFollowHandler(follow_add_follow.NewFollowRequest{
			Context:  context.Background(),
			UserId:   test.UserId,
			ToFollow: test.ToFollow,
		})
		assert.NoError(t, err, "Error calling handler")

		assert.Equal(t, follow_add_follow.NewFollowResponse{
			FollowId: test.ToFollow,
		}, res, "Response incorrect")
	}
}
