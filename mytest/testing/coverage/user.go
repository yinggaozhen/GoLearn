package unittest

type User struct {
	Name string
	Age int
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetAge() int {
	if u.Age == 18 {
		return 18
	}

	if u.Age == 19 {
		return 19
	}

	return u.Age
}