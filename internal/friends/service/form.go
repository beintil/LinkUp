package service

import (
	cs "LinkUp_Update/constans"
	"html/template"
	"net/http"
)

type userData struct {
	UsersData []users
	ID        string
}

func (s *Service) friendsFormHandler(usersData *[]users) {
	var file string

	if s.c.Request.URL.Path != cs.UrlWithoutId(cs.GetFriends, s.c.Request.URL.Path) {
		file = "./internal/friends/html/friends.html"
	} else {
		file = "./internal/friends/html/friends_visited.html"
	}
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}

	if usersData != nil {
		err = tmpl.Execute(s.c.Writer, userData{
			UsersData: *usersData,
			ID:        s.c.Param("id"),
		})
	} else {
		err = tmpl.Execute(s.c.Writer, userData{
			ID: s.c.Param("id"),
		})
	}

	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
	}
}
