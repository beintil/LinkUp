package mycookies

import (
	"LinkUp_Update/config"
	cs "LinkUp_Update/constans"
	"LinkUp_Update/var/logs"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	cookieName = "COOKIE_NAME"
	secretKey1 = "SECRET_KEY_1"
	secretKey2 = "SECRET_KEY_2"
)

// DecodeIdFromCookie This function decodes a cookie and returns the decoded value as a string.
// It takes in the HTTP response writer and request as arguments.
func DecodeIdFromCookie(c *gin.Context) string {
	k1, k2 := getKey()
	s := newService(c, config.Get(cookieName).ToString(), k1, k2)
	id, err := s.decode()
	if err != nil {
		c.Redirect(http.StatusSeeOther, cs.Authorization)
		return id
	}
	return id
}

// Create This function creates a new cookie with the specified ID value.
// It takes in the HTTP response writer, request, and ID string as arguments.
func Create(c *gin.Context, id string) {
	k1, k2 := getKey()
	s := newService(c, config.Get(cookieName).ToString(), k1, k2)
	s.create(id)
}

// CookiesExist This function checks if the cookie exists and is not expired.
// It returns a boolean value and takes in the HTTP response writer and request as arguments.
func CookiesExist(c *gin.Context) bool {
	k1, k2 := getKey()
	s := newService(c, config.Get(cookieName).ToString(), k1, k2)
	return s.check()
}

// Delete This function deletes the cookie by setting its expiration time to the past.
// It takes in the HTTP response writer and request as arguments.
func Delete(c *gin.Context) {
	k1, k2 := getKey()
	s := newService(c, config.Get(cookieName).ToString(), k1, k2)
	s.delete()
}

// getKey This function returns two byte arrays that contain the secret keys used to encrypt and decrypt the cookie.
func getKey() ([]byte, []byte) {
	k1, err := hex.DecodeString(config.Get(secretKey1).ToString())
	if err != nil {
		logs.Get().LogApi(err, func() {
			panic(err)
		})
	}
	k2, err := hex.DecodeString(config.Get(secretKey2).ToString())
	if err != nil {
		logs.Get().LogApi(err, func() {
			panic(err)
		})
	}
	return k1, k2
}
