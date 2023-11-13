package entity

type Photo struct {
	Base
	Title    string `gorm:"not null; type:varchar(255)"`
	Caption  string `gorm:"not null; type:varchar(255)"`
	PhotoURL string `gorm:"not null; type:varchar(255)"`
	UserID   uint   `gorm:"not null; type:int"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE"`
}
