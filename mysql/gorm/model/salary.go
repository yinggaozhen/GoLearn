package model

type Salary struct {
	Id int `gorm:"column:id"`
	Salary int `gorm:"column:salary"`
	People string `gorm:"column:people"`
}
