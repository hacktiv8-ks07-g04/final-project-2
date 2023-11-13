package entity

type SocialMedia struct {
	Base
	Name           string `gorm:"not null; type:varchar(255)"`
	SocialMediaURL string `gorm:"not null; type:varchar(255)"`
	UserID         uint   `gorm:"not null; type:int"`
	User           User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}
