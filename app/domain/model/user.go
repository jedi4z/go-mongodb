package model

type User struct {
	id        string
	firstName string
	lastName  string
	email     string
}

func NewUser(id, firstName, lastName, email string) *User {
	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}
}

func (u *User) GetID() string {
	return u.id
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
