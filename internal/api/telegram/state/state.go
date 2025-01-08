package state

const (
	CreateUserWaiting = "create_user_waiting"
	UpdateUserWaiting = "update_user_waiting"
	AddWishWait       = "add_wish_wait"
	AddWishDesc       = "add_wish_desc"
	AddWishLink       = "add_wish_link"
	AddWishStat       = "add_wish_stat"
	GetUserWish       = "get_user_wish"
	DeleteWish        = "delete_wish"
	AddFriendWait     = "add_friend_wait"
	RemoveFriendWait  = "remove_friend_wait"
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
