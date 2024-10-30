package audio

import (
	"animal-sound-recognizer/internal/db"
	"animal-sound-recognizer/internal/projects"
)

var connection = db.CreateConnection()

func Create(audio AudioEntity) (AudioEntity, error) {
	err := connection.AutoMigrate(AudioDal{})
	if err != nil {
		return AudioEntity{}, err
	}

	var projectDal projects.ProjectDal
	if err := connection.First(&projectDal, audio.ProjectID).Error; err != nil {
		return AudioEntity{}, err
	}

	audioDal := FromEntityToDal(audio)
	connection.Create(&audioDal)

	return FromDalToEntity(audioDal), nil
}

func GetAudio(id uint64) (AudioEntity, error) {
	err := connection.AutoMigrate(AudioDal{})
	if err != nil {
		return AudioEntity{}, err
	}

	var audioDal AudioDal

	result := connection.First(&audioDal, id)
	if result.Error != nil {
		return AudioEntity{}, result.Error
	}

	return FromDalToEntity(audioDal), nil
}

func GetAudios(projectId uint64, limit int, offset int) (AudioEntityList, error) {
	var audioDals []AudioDal

	all := connection.Where("project_id = ?", projectId).Order("created_at DESC")
	all.Limit(limit).Offset(offset).Find(&audioDals)

	var audioEntities = make([]AudioEntity, 0)
	for _, audioDal := range audioDals {
		audioEntities = append(audioEntities, FromDalToEntity(audioDal))
	}

	var total int64
	all.Count(&total)

	return AudioEntityList{
		Audios: audioEntities,
		Total:  total,
	}, nil
}

func EditAudio(id uint64, data EditAudioRequest) error {
	var audioDal AudioDal
	result := connection.First(&audioDal, id)
	if result.Error != nil {
		return result.Error
	}

	audioDal.Name = data.Name
	saveResult := connection.Save(&audioDal)

	if saveResult.Error != nil {
		return saveResult.Error
	}

	return nil
}

func DeleteAudio(id uint64) error {
	var audioDal AudioDal

	result := connection.First(&audioDal, id)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := connection.Delete(&audioDal)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
