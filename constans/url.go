package constans

const (
	Registration  = "/register"
	Authorization = "/auth"
	Home          = "/home"
	HomeWithId    = "/home/:id"
	Profile       = "/home/profile"
	Logout        = "/logout"
	SearchUser    = "/search"
	GetFriends    = "/friends"
	AddFriends    = "/friends/add"
	DeleteFriends = "/friends/delete"
	Chats         = "/chats"
)

func UrlWithoutId(url string, str string) string {
	return url + str[len(url):]
}
