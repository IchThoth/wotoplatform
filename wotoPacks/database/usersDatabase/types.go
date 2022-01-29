package usersDatabase

import wv "wp-server/wotoPacks/core/wotoValues"

type NewUserData struct {
	Username   string
	Password   string
	TelegramId int64
	FirstName  string
	LastName   string
	Birthday   string
	Permission wv.UserPermission
	By         wv.PublicUserId
}
