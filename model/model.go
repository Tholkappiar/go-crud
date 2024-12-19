package model

type Blog struct {
	Id uint `gorm:"primaryKey;autoIncrement`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
}