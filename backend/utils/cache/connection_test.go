package cache_test

import (
	"testing"

	"github.com/nicksan222/ketoai/utils/cache"
)

func TestGetCacheClient(t *testing.T) {
	conn, err := cache.GetCacheClient()

	if err != nil {
		t.Error(err)
	}

	if conn == nil {
		t.Error("Connection is nil")
	}
}
