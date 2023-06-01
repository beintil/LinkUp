package services

import (
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/pkg/decoder"
	"LinkUp_Update/pkg/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"math/rand"
	"net/http"
	"time"
)

func (s *Service) Register() {
	var user User
	if err := s.c.ShouldBindJSON(&user); err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}

	err := validate(user)
	if err != nil {
		s.c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if s.checkEmailExists(&user) {
		s.c.JSON(http.StatusBadRequest, gin.H{
			"message": "this email address exists",
		})
		return
	}

	user.HashPassword, err = decoder.Encode(user.Password)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	err = s.insertUser(&user)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}

	mycookies.Create(s.c, user.ID.String())
	s.c.Status(http.StatusOK)
}

func validate(user User) error {
	var v *validator.Validator
	err := v.Email(user.Email)
	if err != nil {
		return err
	}
	err = v.Password(user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) checkEmailExists(user *User) bool {
	var cont int
	_ = s.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&cont)
	return cont > 0
}

func (s *Service) generateLocalId() int {
	for {
		rand.Seed(time.Now().UnixNano())
		id := rand.Intn(1e9)
		var exists bool
		err := s.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE local_id = $1)", id)
		if err != nil || !exists {
			return id
		}
	}
}

func (s *Service) insertUser(user *User) error {
	var err error
	user.ID, err = uuid.NewV1()
	if err != nil {
		return fmt.Errorf("error generating UUID: %v", err)
	}

	sql := `INSERT INTO users (login, email, password, id, local_id) VALUES ($1, $2, $3, $4, $5)`
	_, err = s.db.Exec(sql, user.Login, user.Email, user.HashPassword, user.ID, s.generateLocalId())
	if err != nil {
		return fmt.Errorf("error inserting user: %v", err)
	}
	return nil
}
