package unittest

type User struct {
	Name string
	Sex int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) IsBoy() bool {
	return u.Sex == 1
}