package audio

import (
	"animal-sound-recognizer/internal/db"
	"animal-sound-recognizer/internal/projects"
)

var connection = db.CreateConnection()

func Create(audio Audio) (Audio, error) {
	err := connection.AutoMigrate(Audio{})
	if err != nil {
		return Audio{}, err
	}

	var project projects.Project
	if err := connection.First(&project, audio.ProjectID).Error; err != nil {
		return Audio{}, err
	}

	connection.Create(&audio)

	return audio, nil
}

func GetAudio(id uint64) (Audio, error) {
	err := connection.AutoMigrate(Audio{})
	if err != nil {
		return Audio{}, err
	}

	var audio Audio

	result := connection.First(&audio, id)
	if result.Error != nil {
		return Audio{}, result.Error
	}

	return audio, nil
}

func GetAudios(projectId uint64, limit int, offset int) (AudioList, error) {
	var audios []Audio

	all := connection.Where("project_id = ?", projectId).Order("created_at DESC")
	all.Limit(limit).Offset(offset).Find(&audios)

	var audioEntities = make([]Audio, 0)
	for _, audio := range audios {
		audioEntities = append(audioEntities, audio)
	}

	var total int64
	all.Count(&total)

	return AudioList{
		Audios: audioEntities,
		Total:  total,
	}, nil
}

func EditAudio(id uint64, data EditAudioRequest) error {
	var audio Audio
	result := connection.First(&audio, id)
	if result.Error != nil {
		return result.Error
	}

	audio.Name = data.Name
	saveResult := connection.Save(&audio)

	if saveResult.Error != nil {
		return saveResult.Error
	}

	return nil
}

func DeleteAudio(id uint64) error {
	var audio Audio

	result := connection.First(&audio, id)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := connection.Delete(&audio)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
