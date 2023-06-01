package routing

import (
	cs "LinkUp_Update/constans"
	"LinkUp_Update/internal/friends"
	"LinkUp_Update/internal/identification"
	"LinkUp_Update/internal/profile"
	"LinkUp_Update/internal/routing/middleware"
	"LinkUp_Update/internal/search"
	"LinkUp_Update/internal/user/home"
	"LinkUp_Update/var/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

type routers struct {
	eng *gin.Engine
	l   *logs.Logging
}

func (r *routers) identification() {
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Conversion(cs.Registration), identification.Registration)
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Conversion(cs.Authorization), identification.Authorization)
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Conversion(cs.Logout), identification.Logout)
}

func (r *routers) profile() {
	var mw = middleware.Middleware{}
	g := r.eng.Use(mw.CheckCookieAndRedirect())
	g.GET(cs.Conversion(cs.Profile), profile.Get)
	g.POST(cs.Conversion(cs.Profile), profile.Edit)
}

func (r *routers) friends() {
	var mw = middleware.Middleware{}
	g := r.eng.Use(mw.CheckCookieAndRedirect())
	g.GET(cs.Conversion(cs.GetFriends), friends.Get)
	g.POST(cs.Conversion(cs.DeleteFriends), friends.Delete)
	g.POST(cs.Conversion(cs.AddFriends), friends.Add)
}

func (r *routers) search() {
	var mw = middleware.Middleware{}
	g := r.eng.Use(mw.CheckCookieAndRedirect())
	g.Match([]string{http.MethodPost, http.MethodGet}, cs.Conversion(cs.SearchUser), search.Search)
}

func (r *routers) user() {
	var mw = middleware.Middleware{}

	r.eng.Use(mw.CheckCookieAndRedirect()).GET(cs.Conversion(cs.Home), home.FormHandler)
}
