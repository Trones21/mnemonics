package user

import "time"

type User struct {
	UserID        string
	NickName      string
	DateJoined    time.Time
	EmailAddress  string
	MnemonicCount int
}
