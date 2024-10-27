package projects

import (
	"time"
)

type ProjectDal struct {
	ID          uint64    `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	//Audios      []audio.AudioDal `gorm:"foreignKey:ID"`
}
