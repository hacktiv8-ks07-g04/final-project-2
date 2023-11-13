package entity

type Comment struct {
	Base
	UserID  uint   `gorm:"not null; type:int"`
	PhotoID uint   `gorm:"not null; type:int"`
	Message string `gorm:"not null; type:varchar(255)"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
	Photo   Photo  `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE"`
}
