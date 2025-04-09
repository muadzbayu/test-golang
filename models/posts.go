package models

import "time"

type Posts struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"type:varchar(200);not null"`
	Content     string    `gorm:"type:text"`
	Category    string    `gorm:"type:varchar(100)"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Status      string    `gorm:"type:varchar(100)"`
}
