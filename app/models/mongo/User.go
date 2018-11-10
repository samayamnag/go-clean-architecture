package mongo

import (
	"time"
	"strings"

	"github.com/samayamnag/boilerplate/app/util"
)

type User struct {
	ID        ID        `bson:"_id" form:"-" json:"id"`
	Admin     bool      `bson:"admin" form:"admin" json:"admin"`
	FullName  string    `bson:"full_name" form:"full_name"`
	Email     string    `bson:"email" form:"email"`
	Password  string    `bson:"password" form:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Users []User

func (user *User) BeforeInsert() error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

func (user *User) BeforeUpdate() error {
	user.UpdatedAt = time.Now()
	return nil
}

func (user *User) SetPasswordField() error {
	user.Password = util.HashAndSalt(user.Password)
	return nil
}

func (user *User) SetEmailField() error {
	user.Email = strings.ToLower(user.Email)
	return nil
}

func(user *User) SetFullNameField() error {
	user.FullName = strings.Title(user.FullName)
	return nil
}
