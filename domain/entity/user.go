package entity

type User struct {
	Base
	Username string `gorm:"not null; type:varchar(255)"`
	Email    string `gorm:"not null; type:varchar(255)"`
	Password string `gorm:"not null; type:varchar(255)"`
	Age      uint   `gorm:"not null; type:int"`
}
