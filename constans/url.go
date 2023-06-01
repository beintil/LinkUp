package constans

const (
	Registration     = "/register"
	Authorization    = "/auth"
	Home             = "/home"
	HomeWithId       = "/home/:id"
	Profile          = "/home/profile"
	ProfileWithId    = "/home/profile/:id"
	Logout           = "/logout"
	SearchUser       = "/search"
	GetFriends       = "/friends"
	GetFriendsWithId = "/friends/:id"
	AddFriends       = "/friends/add"
	DeleteFriends    = "/friends/delete"
	Chats            = "/chats"
)

func UrlWithoutId(url string, str string) string {
	return url + str[len(url):]
}
