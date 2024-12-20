package model

type Blog struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
}

type User struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}