package mycookies

import (
	"github.com/Gorilla/securecookie"
	"github.com/gin-gonic/gin"
	"time"
)

type cookieService interface {
	check() bool
	create(id string)
	delete()
	decode() (string, error)
}

type service struct {
	c          *gin.Context
	cookieName string
	cookie     *securecookie.SecureCookie
}

// newService This function creates a new cookie service with the specified
// c *gin.Context, cookie name, and encryption keys as arguments.
func newService(c *gin.Context, cookieName string, hashKey, blockKey []byte) cookieService {
	return &service{
		c:          c,
		cookieName: cookieName,
		cookie:     securecookie.New(hashKey, blockKey),
	}
}

// decode This method decodes the cookie and returns its value as a string. It takes no arguments.
func (s *service) decode() (string, error) {
	var id string

	cookie, err := s.c.Cookie(s.cookieName)
	if err != nil {
		return id, err
	}

	err = securecookie.DecodeMulti(s.cookieName, cookie, &id, s.cookie)
	return id, err
}

// check This method checks if the cookie exists and is not expired.
// It returns a boolean value and takes no arguments.
func (s *service) check() bool {
	cookie, err := s.c.Cookie(s.cookieName)
	if err != nil || cookie == "" {
		return false
	}
	return true
}

// create This method creates a new cookie with the specified value and expiration time.
// It takes the ID string as an argument.
func (s *service) create(id string) {
	encodedId, err := s.cookie.Encode(s.cookieName, id)
	if err != nil {
		panic(err)
	}

	// Sending cookies to the client via HTTP response
	s.c.SetCookie(s.cookieName, encodedId, int((24 * time.Hour).Seconds()), "/", "localhost", false, false)
}

// delete This method deletes the cookie by setting its expiration time to the past.
// It takes no arguments.
func (s *service) delete() {
	s.c.SetCookie(s.cookieName, "", -1, "/", "localhost", false, false)
}
