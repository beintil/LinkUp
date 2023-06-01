package services

import (
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/pkg/decoder"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (s *Service) Auth() {
	var user User
	err := s.c.ShouldBindJSON(&user)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	user.HashPassword, err = decoder.Encode(user.Password)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	id, err := s.checkEmailAndPasswordExists(&user)
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	mycookies.Create(s.c, id)
	s.c.Status(http.StatusOK)
}

func (s *Service) checkEmailAndPasswordExists(user *User) (string, error) {
	var id string
	err := s.db.QueryRow(`SELECT password, id FROM users WHERE email = $1`, user.Email).Scan(&user.HashPassword, &id)
	if err != nil {
		return id, errors.New("There is no user with such an email")
	}

	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(user.Password))
	if err != nil {
		return id, errors.New("Invalid password")
	}
	return id, nil
}
