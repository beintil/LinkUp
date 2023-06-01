package constans

import (
	"fmt"
	"strings"
)

type Url string

const (
	Registration  Url = "/register"
	Authorization Url = "/auth"
	Home          Url = "/home"
	Profile       Url = "/home/profile"
	Logout        Url = "/logout"
	SearchUser    Url = "/search"
	GetFriends    Url = "/friends"
	AddFriends    Url = "/friends/add"
	DeleteFriends Url = "/friends/delete"
	Chats         Url = "/chats"
)

const (
	HomeWithId       Url = "/home/{id:[^/]+}"
	ProfileWithId    Url = "/home/profile/{id:[^/]+}"
	GetFriendsWithId Url = "/friends/{id:[^/]+}"
	ChatWithFriendId Url = "/chat/{chat_id:[^/]+}"
	Message          Url = "/chat/{chat_id:[^/]+}/message"
)

func UrlWithoutId(url Url, str string) Url {
	url = Url(fmt.Sprint(strings.ReplaceAll(Conversion(url), "{id:[^/]+}", ""), str))
	_ = url
	return url
}
