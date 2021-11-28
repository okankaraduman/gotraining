package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

// Methods

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) SetUserName(uname string) {
	u.UserName = uname
}
