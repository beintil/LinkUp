package service

import (
	mycookies "LinkUp_Update/internal/cookie"
	"errors"
	"net/http"
	"strings"
)

func (s *Service) Search() {
	var user search
	user.Data = s.c.Request.FormValue("data")
	err := validData(&user)
	if err != nil {
		DataFormHandler(s.c, nil)
		return
	}
	result, err := s.searchUser(getData(user), mycookies.DecodeIdFromCookie(s.c))
	if err != nil {
		DataFormHandler(s.c, nil)
		return
	}
	DataFormHandler(s.c, &result)
}

func (s *Service) searchUser(user *search, id string) ([]search, error) {
	var results []searchResult
	sql := `SELECT login, first_name, last_name, local_id 
FROM users 
WHERE (login ILIKE '%' || $1 || '%' OR first_name || ' ' || last_name ILIKE '%' || $1 || '%') 
AND (first_name ILIKE '%' || $2 || '%' OR $2 = '') 
AND (last_name ILIKE '%' || $3 || '%' OR $3 = '')
AND id != $4
ORDER BY 
  CASE 
    WHEN first_name ILIKE '%' || $1 || '%' THEN 1 
    WHEN first_name || ' ' || last_name ILIKE '%' || $1 || '%' THEN 2
    WHEN login ILIKE '%' || $1 || '%' THEN 3
    WHEN last_name ILIKE '%' || $1 || '%' THEN 4
    ELSE 5
  END;
`
	err := s.db.Select(&results, sql, user.Login, user.FirstName, user.LastName, id)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return nil, err
	}

	var myLocalId []any
	err = s.db.Select(&myLocalId, `SELECT local_id FROM users WHERE id = $1`, id)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return nil, err
	}

	var users = make([]search, len(results))
	for n, r := range results {
		users[n].Login = r.Login
		users[n].LocalId = r.LocalId
		users[n].FirstName = r.FirstName.String
		users[n].LastName = r.LastName.String
		var exists bool
		err = s.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE local_id = $1 AND $2 = ANY(friends_id))", myLocalId[0], r.LocalId)
		if err != nil {
			s.l.LogApi(err)
			s.c.Status(http.StatusInternalServerError)
			return nil, err
		}
		users[n].IsFriends = exists
	}
	return users, nil
}

func validData(user *search) error {
	if strings.TrimSpace(user.Data) == "" {
		return errors.New("the fields are empty")
	}
	return nil
}

func getData(user search) *search {
	worlds := strings.Split(user.Data, " ")

	user.Login = worlds[0]
	if len(worlds) >= 2 {
		user.FirstName = worlds[1]
	}
	if len(worlds) >= 3 {
		user.LastName = worlds[2]
	}
	return &user
}
