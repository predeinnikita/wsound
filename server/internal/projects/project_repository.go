package projects

import (
	"animal-sound-recognizer/internal/db"
	"time"
)

var connection = db.CreateConnection()

type ProjectDal struct {
	ID          uint64    `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func Create(project ProjectEntity) (ProjectEntity, error) {
	err := connection.AutoMigrate(ProjectDal{})
	if err != nil {
		return ProjectEntity{}, err
	}

	projectDal := FromEntityToDal(project)
	connection.Create(&projectDal)

	return FromDalToEntity(projectDal), nil
}

func GetProject(id uint64) (ProjectEntity, error) {
	err := connection.AutoMigrate(ProjectDal{})
	if err != nil {
		return ProjectEntity{}, err
	}

	var projectDal ProjectDal

	result := connection.First(&projectDal, id)
	if result.Error != nil {
		return ProjectEntity{}, result.Error
	}

	return FromDalToEntity(projectDal), nil
}

func GetProjects(limit int, offset int) (ProjectEntityList, error) {
	var projectDals []ProjectDal

	connection.Order("created_at DESC").Limit(limit).Offset(offset).Find(&projectDals)

	var projectEntities = make([]ProjectEntity, 0)
	for _, projectDal := range projectDals {
		projectEntities = append(projectEntities, FromDalToEntity(projectDal))
	}

	return ProjectEntityList{
		Projects: projectEntities,
	}, nil
}

func Update(id uint64, project ProjectEntity) error {
	var projectDal ProjectDal
	result := connection.First(&projectDal, id)
	if result.Error != nil {
		return result.Error
	}

	projectDal.Name = project.Name
	projectDal.Description = project.Description
	saveResult := connection.Save(&projectDal)

	if saveResult.Error != nil {
		return saveResult.Error
	}

	return nil
}

func Delete(id uint64) error {
	var projectDal ProjectDal

	result := connection.First(&projectDal, id)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := connection.Delete(&projectDal)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
