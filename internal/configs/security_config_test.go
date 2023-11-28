package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCheckTokenValidity_ValidToken(t *testing.T) {
	// Given
	introspectionResponse := map[string]interface{}{
		"active": true,
	}

	c := &gin.Context{}

	// When
	err := checkTokenValidity(introspectionResponse, c)

	// Then
	assert.NoError(t, err)
}

func TestCheckTokenValidity_InvalidToken(t *testing.T) {
	// Given
	introspectionResponse := map[string]interface{}{
		"active": false,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	// When
	err := checkTokenValidity(introspectionResponse, c)

	// Then
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid token")
}

func TestCheckRolesAccess(t *testing.T) {
	// Given
	requiredRoles := []string{"role1", "role2"}

	introspectionResponse := map[string]interface{}{
		"realm_access": map[string]interface{}{
			"roles": []interface{}{"role1", "role2"},
		},
	}
	c := &gin.Context{}

	// When
	err := checkRolesAccess(introspectionResponse, requiredRoles, c)

	// Then
	assert.NoError(t, err)
}
