package http

import "students/pkg/user"

func (c Controller) toUserDto(u User) *user.User {
	return user.NewUser(u.Name, u.Age)
}

func (c Controller) toResponseItem(u *user.User) *User {
	return &User{
		Id:   u.Id,
		Name: u.Name,
		Age:  u.Age,
	}
}
