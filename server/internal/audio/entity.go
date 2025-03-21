package audio

import "time"

type Audio struct {
	ID        uint64     `json:"id"         gorm:"primaryKey"`
	StorageID string     `json:"storage_id" gorm:"size:255;not null"`
	Name      string     `json:"name"       gorm:"size:255;not null"`
	ProjectID uint64     `json:"project_id" gorm:"size:255;not null"`
	Status    Status     `json:"status"     gorm:"default:not_wolf"`
	Intervals []Interval `json:"intervals"  gorm:"foreignKey:AudioID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

type Interval struct {
	ID      uint64 `json:"id"       gorm:"primaryKey"`
	AudioID uint64 `json:"audio_id" gorm:"not null;index"`
	Start   string `json:"start"    gorm:"not null"`
	End     string `json:"end"      gorm:"not null"`
}

type AudioList struct {
	Audios []Audio `json:"audios"`
	Total  int64   `json:"total"`
}

type EditAudioRequest struct {
	Name string `json:"name"`
}
