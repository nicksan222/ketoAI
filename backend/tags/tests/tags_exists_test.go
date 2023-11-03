package tags

import (
	"testing"

	"github.com/nicksan222/ketoai/tags"
	"github.com/stretchr/testify/assert"
)

func TestTagsExists(t *testing.T) {
	var tagsIds []string = []string{}

	// 1. Create Many tags
	for i := 0; i < 10; i++ {
		inputData := `{"name":"Test Tag"}`
		request, err := tags.ParseCreateTagRequest([]byte(inputData))
		assert.NoError(t, err, "Failed to parse create tag request")

		// Call the CreateTag function
		response, err := tags.CreateTag(request)
		assert.NoError(t, err, "Failed to create tag")
		assert.NotEmpty(t, response.TagId, "No tag ID returned after creation")

		tagsIds = append(tagsIds, response.TagId)
	}

	// 2. Check if tags exists
	tagsExistsStringRequest := ""

	for i := 0; i < 10; i++ {
		tagsExistsStringRequest += tagsIds[i] + ","
	}

	tagsExistsRequest := tags.TagsExistsRequest{
		TagIds: tagsIds,
	}

	tagsExistsResponse, err := tags.TagsExists(tagsExistsRequest)

	assert.NoError(t, err, "Failed to check if tags exists")
	assert.NotEmpty(t, tagsExistsResponse.Exists, "No tags exists")
	assert.Len(t, tagsExistsResponse.Exists, 10, "Tags exists")
	assert.Len(t, tagsExistsResponse.NotExists, 0, "Tags exists")

	// Searching a non-existing tag
	tagsExistsRequest = tags.TagsExistsRequest{
		TagIds: []string{"5f9b3b3b3b3b3b3b3b3b3b3b"},
	}

	tagsExistsResponse, err = tags.TagsExists(tagsExistsRequest)

	assert.NoError(t, err, "Failed to check if tags exists")
	// assert.Equal(t, tagsExistsResponse.Exists, "[]", "Tags exists")
	// assert.Len(t, tagsExistsResponse.NotExists, 1, "Tags exists")

	// 3. Delete the created tags
	for i := 0; i < 10; i++ {
		deleteRequest := tags.DeleteTagRequest{
			TagId: tagsIds[i],
		}
		deleteResponse, err := tags.DeleteTag(deleteRequest)
		assert.NoError(t, err, "Failed to delete Tag")
		assert.True(t, deleteResponse.Deleted, "Tag not deleted")
	}
}
