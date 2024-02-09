package models

import (
	"time"
)

type Session struct {
	SessionID   int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt   time.Time
	SessionName string
	OwnerID     int64  // This is the foreign key for the owner
	Owner       User   `gorm:"foreignKey:OwnerID"`       // Define the owner relationship
	Users       []User `gorm:"many2many:session_users;"` // Many-to-many relationship with users
}
