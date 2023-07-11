package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestInitStore(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}
