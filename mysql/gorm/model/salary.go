package model

// - 模型定义 : {@link https://gorm.io/docs/models.html}
type Salary struct {
	Id int `gorm:"column:id"`
	Salary int `gorm:"column:salary"`
	People string `gorm:"column:people"`
}
