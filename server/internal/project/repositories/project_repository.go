package project_repository

import (
	"animal-sound-recognizer/internal/db"
	"animal-sound-recognizer/internal/project/entities"
	"time"
)

var connection = db.CreateConnection()

type ProjectDal struct {
	ID          uint64    `gorm:"primaryKey"`
	Name        string    `gorm:"size:255;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func projectDalToProjectEntity(dal ProjectDal) project_entities.ProjectEntity {
	return project_entities.ProjectEntity{
		ID:          dal.ID,
		Name:        dal.Name,
		Description: dal.Description,
		CreatedAt:   dal.CreatedAt,
	}
}

func projectEntityToProjectDal(entity project_entities.ProjectEntity) ProjectDal {
	return ProjectDal{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
	}
}

func Create(project project_entities.ProjectEntity) (project_entities.ProjectEntity, error) {
	err := connection.AutoMigrate(ProjectDal{})
	if err != nil {
		return project_entities.ProjectEntity{}, err
	}

	projectDal := projectEntityToProjectDal(project)
	connection.Create(&projectDal)

	return projectDalToProjectEntity(projectDal), nil
}

func GetProject(id uint64) (project_entities.ProjectEntity, error) {
	err := connection.AutoMigrate(ProjectDal{})
	if err != nil {
		return project_entities.ProjectEntity{}, err
	}

	var projectDal ProjectDal

	result := connection.First(&projectDal, id)
	if result.Error != nil {
		return project_entities.ProjectEntity{}, result.Error
	}

	return projectDalToProjectEntity(projectDal), nil
}

func GetProjects(limit int, offset int) (project_entities.ProjectEntityList, error) {
	var projectDals []ProjectDal

	connection.Order("created_at DESC").Limit(limit).Offset(offset).Find(&projectDals)

	var projectEntities = make([]project_entities.ProjectEntity, 0)
	for _, projectDal := range projectDals {
		projectEntities = append(projectEntities, projectDalToProjectEntity(projectDal))
	}

	return project_entities.ProjectEntityList{
		Projects: projectEntities,
	}, nil
}

func Update(id uint64, project project_entities.ProjectEntity) error {
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
