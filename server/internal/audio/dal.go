package audio

import "time"

type AudioDal struct {
	ID        uint64    `gorm:"primaryKey"`
	Name      string    `gorm:"size:255;not null"`
	StorageID string    `gorm:"size:255;not null"`
	ProjectID uint64    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
