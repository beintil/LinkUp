package service

import (
	cs "LinkUp_Update/constans"
	mycookies "LinkUp_Update/internal/cookie"
	"net/http"
	"strconv"
	"time"
)

func (s *Service) Edit() {
	user, err := s.parseEditData()
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	err = s.updateUser(user)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	s.c.Redirect(http.StatusFound, cs.Profile)
}

func (s *Service) parseEditData() (editUserData, error) {
	var r = s.c.Request
	err := r.ParseForm()
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return editUserData{}, err
	}

	user := editUserData{
		ID:          mycookies.DecodeIdFromCookie(s.c),
		Email:       r.FormValue("email"),
		Login:       r.FormValue("login"),
		FirstName:   r.FormValue("firstName"),
		LastName:    r.FormValue("lastName"),
		Gender:      r.FormValue("gender"),
		DateOfBirth: r.FormValue("date_of_birth"),
	}

	user.Age, err = serviceAge(r.FormValue("age"))
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return user, err
	}

	return user, nil
}

func serviceAge(str string) (int, error) {
	if str != "" {
		age, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		return age, nil
	}
	return 0, nil
}

func (s *Service) updateUser(user editUserData) error {
	sql := `UPDATE users SET `
	sqlParams := make(map[string]interface{})

	if user.Login != "" {
		sql += "login = :login, "
		sqlParams["login"] = user.Login
	}
	if user.Email != "" {
		sql += "email = :email, "
		sqlParams["email"] = user.Email
	}
	if user.Gender != "" {
		sql += "gender = :gender, "
		sqlParams["gender"] = user.Gender
	}
	if user.DateOfBirth != "" {
		sql += "date_of_birth = :date_of_birth, "
		sqlParams["date_of_birth"] = user.DateOfBirth
	}
	if user.LastName != "" {
		sql += "last_name = :last_name, "
		sqlParams["last_name"] = user.LastName
	}
	if user.FirstName != "" {
		sql += "first_name = :first_name, "
		sqlParams["first_name"] = user.FirstName
	}
	if user.Age != 0 {
		sql += "age = :age, "
		sqlParams["age"] = user.Age
	}

	// Remove the trailing comma and space from the SQL string
	sql = sql[:len(sql)-2]

	// Add the WHERE clause
	sql += ` WHERE id = :id`

	// Add the id parameter to the sqlParams map
	sqlParams["id"] = user.ID

	_, err := s.db.NamedExec(sql, sqlParams)
	if err != nil {
		return err
	}

	return nil
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
