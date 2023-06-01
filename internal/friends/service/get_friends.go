package service

import (
	mycookies "LinkUp_Update/internal/cookie"
	"github.com/jmoiron/sqlx"
	"html/template"
	"net/http"
	"strings"
)

func (s *Service) Get() {
	var myFriends []string
	var err error

	myFriends, err = s.getFriendsIdFromDb("SELECT friends_id FROM users WHERE id = $1", mycookies.DecodeIdFromCookie(s.c))

	if err != nil {
		s.l.LogApi(err)
		return
	}
	if len(myFriends) > 0 {
		if myFriends[0] == "{}" {
			s.friendsFormHandler(nil)
			return
		}

		data, err := s.getUserData(myFriends)
		if err != nil {
			return
		}
		var usersData = make([]users, len(data))
		for n, v := range data {
			usersData[n].Login = v.Login
			usersData[n].LocalId = v.LocalId
			usersData[n].FirstName = v.FirstName.String
			usersData[n].LastName = v.LastName.String
		}
		s.friendsFormHandler(&usersData)
		return
	}
	s.friendsFormHandler(nil)
}

func (s *Service) getUserData(friends []string) ([]dataFromDb, error) {
	idList := strings.Split(friends[0][1:len(friends[0])-1], ",")

	query := `SELECT login, first_name, last_name, local_id FROM users WHERE local_id IN (?)`
	placeholders, args, err := sqlx.In(query, idList)
	if err != nil {
		return nil, err
	}

	placeholders = s.db.Rebind(placeholders)
	var data []dataFromDb
	return data, sqlx.Select(s.db, &data, placeholders, args...)
}

func (s *Service) getFriendsIdFromDb(sql string, id string) ([]string, error) {
	var value []any
	err := sqlx.Select(s.db, &value, sql, id)
	if err != nil {
		return nil, err
	}
	var myFriends = make([]string, 0)
	for _, v := range value {
		if v != nil {
			myFriends = append(myFriends, v.(string))
		}
	}
	return myFriends, nil
}

func (s *Service) friendsFormHandler(usersData *[]users) {
	tmpl, err := template.ParseFiles("./internal/friends/html/friends.html")
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	type userData struct {
		UsersData []users
		ID        string
	}

	if usersData != nil {
		err = tmpl.Execute(s.c.Writer, userData{
			UsersData: *usersData,
			ID:        s.c.Param("id"),
		})
		if err != nil {
			s.l.LogApi(err)
			s.c.Status(http.StatusInternalServerError)
		}
	} else {
		err = tmpl.Execute(s.c.Writer, userData{
			ID: s.c.Param("id"),
		})
		if err != nil {
			s.l.LogApi(err)
			s.c.Status(http.StatusInternalServerError)
		}
	}
}
