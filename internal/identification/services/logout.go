package services

import (
	cs "LinkUp_Update/constans"
	mycookies "LinkUp_Update/internal/cookie"
	"net/http"
)

func (s *Service) Logout() {
	mycookies.Delete(s.c)
	s.c.Redirect(http.StatusSeeOther, cs.Authorization)
}
