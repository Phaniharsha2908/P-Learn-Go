package model

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUserNotFound(t *testing.T) {
	user, appErr := UserService.GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, appErr, "We were expecting an error with id 0")

	assert.EqualValues(t, http.StatusNotFound, appErr.StatusCode)
	assert.EqualValues(t, "not_found", appErr.Code)
}
