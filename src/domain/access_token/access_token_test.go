package access_token

import (
	"testing"
	"time"

	"github.com/avalokitasharma/bookstore_oauth_api/src/repository/rest"
	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	repository := rest.NewRestUsersRepository()

	user, err := repository.LoginUser("email@gmail.com", "the-password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	at := GetNewAccessToken(user.Id)
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
