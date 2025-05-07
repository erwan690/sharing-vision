package model

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:200;not null"`
	Content     string    `gorm:"type:text;not null"`
	Category    string    `gorm:"size:100"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Status      string    `gorm:"size:100;check:(status IN ('Publish','Draft','Trash'))"`
}
