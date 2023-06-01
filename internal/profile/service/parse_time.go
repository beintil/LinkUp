package service

import (
	"net/http"
	"time"
)

func (s *Service) parseTime(str string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z", str)
	if err != nil {
		s.l.LogApi(err)
		s.c.Status(http.StatusInternalServerError)
		return ""
	}
	return t.Format("2006-01-02")
}
