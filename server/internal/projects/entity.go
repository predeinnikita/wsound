package projects

import "time"

type ProjectEntity struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProjectEntityList struct {
	Projects []ProjectEntity `json:"projects"`
	Total    int64           `json:"total"`
}
