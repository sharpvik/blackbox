package filters

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sharpvik/blackbox/test_utils"
)

func TestMethods(t *testing.T) {
	methodsFilter := Methods(http.MethodPost)
	assert.True(t, methodsFilter.Accepts(test_utils.Post(t, "/", []byte{})))
	assert.False(t, methodsFilter.Accepts(test_utils.Get(t, "/")))
}
