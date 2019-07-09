package model

import "time"

type User struct {
	id        string
	createdAt time.Time
	firstName string
	lastName  string
	email     string
}

func NewUser(id string, createdAt time.Time, firstName, lastName, email string) *User {
	return &User{
		id:        id,
		createdAt: createdAt,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetFirstName() string {
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetEmail() string {
	return u.email
}
