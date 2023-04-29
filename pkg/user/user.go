package user

import (
	"html"
	"strings"
)

type User struct {
	Id   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
	Age  int    `gorm:"not null" json:"age" json:"age"`
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
}
