package audio

import (
	"animal-sound-recognizer/internal/db"
	"animal-sound-recognizer/internal/projects"
	"fmt"
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

	fmt.Println(projectDal)
	fmt.Println(audio)

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

	connection.Where("project_id = ?", projectId).Order("created_at DESC").Limit(limit).Offset(offset).Find(&audioDals)

	var audioEntities = make([]AudioEntity, 0)
	for _, audioDal := range audioDals {
		audioEntities = append(audioEntities, FromDalToEntity(audioDal))
	}

	return AudioEntityList{
		Audios: audioEntities,
	}, nil
}
