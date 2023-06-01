package service

import (
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/var/logs"
	"net/http"
	"time"
)

func (s *Service) Get() {
	var user getUserData
	err := s.getUserFromDB("SELECT email, login, age, first_name, last_name, gender, date_of_birth FROM users WHERE id = $1", &user, mycookies.DecodeIdFromCookie(s.c))
	if err != nil {
		logs.Get().LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}

	s.myDataFormHandler(&user)
}

func (s *Service) getUserFromDB(sql string, user *getUserData, id string) error {
	row := s.db.QueryRow(sql, id)
	if row.Err() != nil {
		return row.Err()
	}
	var get getFromDb
	if err := row.Scan(&user.Email, &user.Login, &get.AgeDB, &get.FirstNameDB, &get.LastNameDB, &get.GenderDB, &get.DateOfBirthDB); err != nil {
		return err
	}
	if get.DateOfBirthDB.String != "" {
		t, err := time.Parse("2006-01-02T15:04:05Z", get.DateOfBirthDB.String)
		if err != nil {
			return err
		}
		user.DateOfBirth = t.Format("2006-01-02")
	}

	user.Age = get.AgeDB.Int64
	user.FirstName = get.FirstNameDB.String
	user.LastName = get.LastNameDB.String
	user.Gender = get.GenderDB.String
	return nil
}
