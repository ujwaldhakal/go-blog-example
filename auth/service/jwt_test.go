package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJwtTokenGeneratesToken(t *testing.T) {
	token, err := GenerateJwtToken("john@doe.com")
	assert.NotEmpty(t, token)
	assert.Empty(t, err)
}
