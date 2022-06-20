package entity

type UserProperty struct {
	Id           int    `gorm:"primaryKey;column:id"`
	UserId       int    `gorm:"column:user_id"`
	PropertyName string `gorm:"column:property_name"`
	Property     string `gorm:"column:property"`
}

func (userProperty UserProperty) TableName() string {
	return "hibernate_test.user_property"
}
