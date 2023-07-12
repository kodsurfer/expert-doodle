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

func TestSaveAndGet(t *testing.T) {
	original := "https://wiki.mozilla.org/Contribute"
	short := "enviewrbivwpbe"

	SaveUrlMapping(short, original)

	getOriginalUrl := GetOriginalUrl(short)

	assert.Equal(t, original, getOriginalUrl)
}
