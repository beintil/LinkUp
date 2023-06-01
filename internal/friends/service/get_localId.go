package service

import "github.com/jmoiron/sqlx"

func (s *Service) getLocalId(myId string) int64 {
	var myLocalId []any
	err := sqlx.Select(s.db, &myLocalId, `SELECT local_id FROM users WHERE id = $1`, myId)
	if err != nil {
		panic("extraction error local id")
	}

	return myLocalId[0].(int64)
}
