package service

import (
	mycookies "LinkUp_Update/internal/cookie"
	"LinkUp_Update/internal/database"
	"LinkUp_Update/var/logs"
	"context"
	"time"
)

func (s *Service) Get() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var db = database.Init(ctx).GetConn()
	defer func() {
		database.Close(db)
		if rec := recover(); rec != nil {
			logs.Get().LogApi(rec)
		}
	}()

	var user getUserData
	err := s.getUserFromDB("SELECT email, login, age, first_name, last_name, gender, date_of_birth FROM users WHERE id = $1", &user, mycookies.DecodeIdFromCookie(s.c))
	if err != nil {
		logs.Get().LogApi(err)
	}

	s.myDataFormHandler(&user)
}
