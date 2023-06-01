package service

import (
	mycookies "LinkUp_Update/internal/cookie"
	"net/http"
	"strconv"
)

func (s *Service) Delete() {
	id := s.c.Request.FormValue("id")
	s.deleteFriend(mycookies.DecodeIdFromCookie(s.c), id)
}

func (s *Service) deleteFriend(myId, idFriends string) {
	id, err := strconv.ParseInt(idFriends, 10, 32)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return
	}
	_, err = s.db.Exec(`UPDATE users SET friends_id = array_remove(friends_id, $1) WHERE id = $2;`, id, myId)
	if err != nil {
		s.c.Status(http.StatusInternalServerError)
		panic(err)
	}

	_, err = s.db.Exec(`UPDATE users SET friends_id = array_remove(friends_id, $1) WHERE local_id = $2`, s.getLocalId(myId), id)
	if err != nil {
		s.c.Status(http.StatusInternalServerError)
		panic(err)
	}
}
