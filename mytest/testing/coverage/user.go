package unittest

type User struct {
	Name string
	Age int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	return u.Age
}