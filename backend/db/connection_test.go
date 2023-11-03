package db

import (
	"testing"
)

func ConnectionTest(t *testing.T) {
	connection, err := GetDBClient()

	if err != nil {
		t.Error(err)
	}

	if connection == nil {
		t.Error("Connection is nil")
	}

}
