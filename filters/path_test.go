package filters

import (
	"testing"

	"github.com/sharpvik/blackbox/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	pathFilter := Path("/api")
	assert.True(t, pathFilter.Accepts(test_utils.Get(t, "/api")))
	assert.False(t, pathFilter.Accepts(test_utils.Get(t, "/api/user")))
}

func TestPathPrefix(t *testing.T) {
	pathPrefixFilter := PathPrefix("/api")
	assert.True(t, pathPrefixFilter.Accepts(test_utils.Get(t, "/api")))
	assert.True(t, pathPrefixFilter.Accepts(test_utils.Get(t, "/api/public")))
	assert.False(t, pathPrefixFilter.Accepts(test_utils.Get(t, "/user")))
}
