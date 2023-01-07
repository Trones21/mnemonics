package user

import (
	"database/sql"
	"fmt"
	"mnemonics/database"
	"time"

	"github.com/google/uuid"
)

func getUser(userID string) (*User, error) {
	row := database.DbConn.QueryRow(`Select 
	UserID,
	NickName,
	DateJoined,
	EmailAddress,
	MnemonicCount
	from user where UserID = ?`, userID)
	user := &User{}
	err := row.Scan(
		&user.UserID,
		&user.NickName,
		&user.DateJoined,
		&user.EmailAddress,
		&user.MnemonicCount,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func getuserList() ([]User, error) {
	results, err := database.DbConn.Query(`select * from users`)
	if err != nil {
		fmt.Println("Err getting rows")
		return nil, err

	}
	defer results.Close()
	users := make([]User, 0)
	for results.Next() {
		var user User
		results.Scan(
			&user.UserID,
			&user.NickName,
			&user.DateJoined,
			&user.MnemonicCount)
		users = append(users, user)
	}
	fmt.Println(users[0].UserID)
	return users, nil

}

func addUser(user *User) error {
	var err error
	_, err = database.DbConn.Exec(`INSERT INTO user 
	(
	UserID,
	NickName,
	DateJoined,
	EmailAddress,
	MnemonicCount
	)
	 VALUES (?,?,?,?,?)`,
		uuid.NewString(),
		user.NickName,
		time.Now(),
		user.EmailAddress,
		0,
	)

	if err != nil {
		return err
	}
	return nil
}

// Note: User.service.go checks that loggedin UserID = UserID from json body, therefore
// ensuring only the loggedin user can update their profile (nobody else can)
func updateUser(user *User) error {
	var err error
	_, err = database.DbConn.Exec(`
	update user
	set NickName = ? ,
	EmailAddress = ?  
	where UserID = ?;
	`,
		user.NickName,
		user.EmailAddress,
		user.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}
