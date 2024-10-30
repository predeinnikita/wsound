package audio

type AudioEntity struct {
	ID        uint64 `json:"id"`
	StorageID string `json:"storage_id"`
	Name      string `json:"name"`
	ProjectID uint64 `json:"project_id"`
}

type AudioEntityList struct {
	Audios []AudioEntity `json:"audios"`
	Total  int64         `json:"total"`
}

type EditAudioRequest struct {
	Name string `json:"name"`
}
