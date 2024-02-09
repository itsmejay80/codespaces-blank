package models

type User struct {
	UserID   int64     `gorm:"primaryKey;autoIncrement"`
	Username string    `gorm:"unique"`
	Sessions []Session `gorm:"many2many:user_sessions;"`
}
