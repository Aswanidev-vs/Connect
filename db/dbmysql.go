package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Aswanidev-vs/Connect/config"
)

type Userinfo struct {
	ID       int
	Username string
	Email    string
	Password string // store hashed password
}

func GetUserByEmail(email string) (*Userinfo, error) {
	row := config.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email=?", email)
	user := &Userinfo{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}

func NewUser(username, email, password string) error {
	if config.DB == nil {
		return fmt.Errorf("database not initialized")
	}

	_, err := config.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, password)
	return err

}

func CheckUser(email string) (bool, error) {
	var id int
	err := config.DB.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // email not found
		}
		return false, err // other errors
	}
	return true, nil
}
