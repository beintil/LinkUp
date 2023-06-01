package service

import (
	"html/template"
)

func (s *Service) myDataFormHandler(user *getUserData) {
	tmpl, err := template.ParseFiles("./internal/profile/html/profile.html")
	if err != nil {
		s.l.LogApi(err)
		return
	}

	err = tmpl.Execute(s.c.Writer, user)
	if err != nil {
		s.l.LogApi(err)
	}
}

func (s *Service) visitedDataFormHandler(user *getUserData) {
	tmpl, err := template.ParseFiles("./internal/profile/html/visit_profile.html")
	if err != nil {
		s.l.LogApi(err)
		return
	}

	err = tmpl.Execute(s.c.Writer, user)
	if err != nil {
		s.l.LogApi(err)
	}
}
