package entity

type User struct {
	Id             int            `gorm:"primaryKey;column:id"`
	FirstName      string         `gorm:"column:first_name"`
	LastName       string         `gorm:"column:last_name"`
	Surname        string         `gorm:"column:surname"`
	UserProperties []UserProperty `gorm:"foreignKey:UserId;references:Id"`
}

func (user User) TableName() string {
	return "hibernate_test.users"
}
