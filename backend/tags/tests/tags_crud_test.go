package tags

import (
	"testing"

	"github.com/nicksan222/ketoai/tags"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteTag(t *testing.T) {
	// 1. Create Tag
	inputData := `{"name":"Test Tag"}`
	request, err := tags.ParseCreateTagRequest([]byte(inputData))
	assert.NoError(t, err, "Failed to parse create tag request")

	// Call the CreateTag function
	response, err := tags.CreateTag(request)
	assert.NoError(t, err, "Failed to create tag")
	assert.NotEmpty(t, response.TagId, "No tag ID returned after creation")

	// 2. Delete the created Tag
	deleteRequest := tags.DeleteTagRequest{
		TagId: response.TagId,
	}
	deleteResponse, err := tags.DeleteTag(deleteRequest)
	assert.NoError(t, err, "Failed to delete Tag")
	assert.True(t, deleteResponse.Deleted, "Tag not deleted")

	// Confirm that the Tag was deleted (you can fetch and make sure it doesn't exist, etc.)
}

// Helper function to clean up any test data
func tearDownTag(TagID string) {
	request := tags.DeleteTagRequest{
		TagId: TagID,
	}
	tags.DeleteTag(request)
}

func TestParseCreateTagRequest(t *testing.T) {
	validJSON := `{"name":"Test Tag","quantity_measurement":"100 grams"}`
	request, err := tags.ParseCreateTagRequest([]byte(validJSON))
	assert.NoError(t, err, "Failed to parse valid JSON")
	assert.Equal(t, "Test Tag", request.Name, "Parsed name doesn't match")

	invalidJSON := `{"name":}`
	_, err = tags.ParseCreateTagRequest([]byte(invalidJSON))
	assert.Error(t, err, "Expected error while parsing invalid JSON")
}

// More tests can follow for other functions
