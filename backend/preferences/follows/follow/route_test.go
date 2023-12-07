package follow_add_follow_test

import (
	"context"
	"testing"

	follow_add_follow "github.com/nicksan222/ketoai/preferences/follows/follow"
	"github.com/stretchr/testify/assert"
)

type NewFollowTest struct {
	UserId      string
	ToFollow    string
	ExpectError bool
	resultCode  int
	res         follow_add_follow.NewFollowResponse
}

var NewFollowTests = []NewFollowTest{
	{
		UserId:     "test_add_follow_user_route",
		ToFollow:   "test_add_follow_user_route",
		resultCode: 400,
		ExpectError: true,
		res: follow_add_follow.NewFollowResponse{
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

		if test.ExpectError {
			assert.Error(t, err, "Expected error but got none")
		} else {
			assert.NoError(t, err, "Unexpected error")
			assert.Equal(t, test.res, res, "Response incorrect")
		}
	}
}
