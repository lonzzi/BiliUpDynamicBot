package db

type Bot struct {
	UID    uint   `gorm:"type:integer;not null;primary_key"`
	Name   string `gorm:"type:varchar(255);not null"`
	Face   string `gorm:"type:text"`
	Cookie string `gorm:"type:text;not null"`
	UserID uint   `gorm:"type:integer;not null"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
