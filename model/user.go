package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
