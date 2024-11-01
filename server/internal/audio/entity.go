package audio

import "time"

type Audio struct {
	ID        uint64    `json:"id"         gorm:"primaryKey"`
	StorageID string    `json:"storage_id" gorm:"size:255;not null"`
	Name      string    `json:"name"       gorm:"size:255;not null"`
	ProjectID uint64    `json:"project_id" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type AudioList struct {
	Audios []Audio `json:"audios"`
	Total  int64   `json:"total"`
}

type EditAudioRequest struct {
	Name string `json:"name"`
}
