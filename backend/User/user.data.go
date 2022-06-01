package user

import (
	"database/sql"
	"fmt"
	"mnemonics/database"
)

func getuser(userID string) (*User, error) {
	row := database.DbConn.QueryRow(`Select 
	userID,
	NickName,
	DateJoined,
	MnemonicCount
	from user where userID = ?`, userID)
	user := &User{}
	err := row.Scan()
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

func addUser(user *User) {

}

func updateUser(user *User) {

}
