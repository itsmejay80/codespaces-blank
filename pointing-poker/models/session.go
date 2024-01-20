package models

import (
	"time"
)


type Session struct {
	ID int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	CreatorID int64
}
