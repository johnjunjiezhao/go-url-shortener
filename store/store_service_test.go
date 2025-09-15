package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StoreService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://example.com"
	shortCode := "exmpl"
	userID := "user123"

	SaveURLMapping(shortCode, initialLink, userID)
	retrievedLink := RetrieveOriginalURL(shortCode)

	assert.Equal(t, initialLink, retrievedLink)
}
