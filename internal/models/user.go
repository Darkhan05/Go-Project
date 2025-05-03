package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;not null"`
}
