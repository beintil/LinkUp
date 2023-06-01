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
	mw  *middleware.Middleware
}

func (r *routers) identification() {
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Registration, identification.Registration)
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Authorization, identification.Authorization)
	r.eng.RouterGroup.Match([]string{http.MethodPost, http.MethodGet}, cs.Logout, identification.Logout)
}

func (r *routers) profile() {
	r.eng.GET(cs.Profile, profile.Get).POST(cs.Profile, profile.Edit)
}

func (r *routers) friends() {
	r.eng.GET(cs.GetFriends, friends.Get)
	r.eng.POST(cs.DeleteFriends, friends.Delete)
	r.eng.POST(cs.AddFriends, friends.Add)
}

func (r *routers) search() {
	r.eng.Match([]string{http.MethodPost, http.MethodGet}, cs.SearchUser, search.Search)
}

func (r *routers) user() {
	r.eng.GET(cs.Home, home.Home)
}

func (r *routers) visit() {
	r.eng.Use(r.mw.HandlerVisitors())
	r.eng.GET(cs.HomeWithId, home.Visit)
	r.eng.GET(cs.ProfileWithId, profile.GetVisit)
	r.eng.GET(cs.GetFriendsWithId, friends.GetVisit)
}
