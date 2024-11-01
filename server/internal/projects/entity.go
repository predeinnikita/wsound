package projects

import "time"

type Project struct {
	ID          uint64    `json:"id"          gorm:"primaryKey"`
	Name        string    `json:"name"        gorm:"size:255;not null"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"  gorm:"autoCreateTime"`
}

type ProjectList struct {
	Projects []Project `json:"projects"`
	Total    int64     `json:"total"`
}
