package project_entities

import "time"

type ProjectEntity struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProjectEntityList struct {
	Projects []ProjectEntity `json:"projects"`
}
