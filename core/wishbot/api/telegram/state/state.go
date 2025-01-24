package state

const (
	CreateUserWaiting = "create_user_waiting"
	CreateUserAdress  = "create_user_adress"
	CreateUserName    = "create_user_name"
	CreateUserPhone   = "create_user_phone"
	UpdateUserWaiting = "update_user_waiting"
	GetUserWish       = "get_user_wish"
	AddFriendWait     = "add_friend_wait"
)

var UserStates = make(map[int64]string)

func SetUserState(chatID int64, usrstate string) {
	UserStates[chatID] = usrstate
}

func GetUserState(chatID int64) string {
	return UserStates[chatID]
}

func ClearUserState(chatID int64) {
	delete(UserStates, chatID)
}
