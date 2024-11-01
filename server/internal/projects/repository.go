package projects

import (
	"animal-sound-recognizer/internal/db"
)

var connection = db.CreateConnection()

func Create(project Project) (Project, error) {
	err := connection.AutoMigrate(Project{})
	if err != nil {
		return Project{}, err
	}

	connection.Create(&project)

	return project, nil
}

func GetProject(id uint64) (Project, error) {
	err := connection.AutoMigrate(Project{})
	if err != nil {
		return Project{}, err
	}

	var project Project

	result := connection.First(&project, id)
	if result.Error != nil {
		return Project{}, result.Error
	}

	return project, nil
}

func GetProjects(limit int, offset int) (ProjectList, error) {
	var projects []Project
	connection.Order("created_at DESC").Limit(limit).Offset(offset).Find(&projects)

	var projectEntities = make([]Project, 0)
	for _, project := range projects {
		projectEntities = append(projectEntities, project)
	}

	var total int64
	connection.Model(&Project{}).Count(&total)

	return ProjectList{
		Projects: projectEntities,
		Total:    total,
	}, nil
}

func Update(id uint64, newProject Project) error {
	var project Project
	result := connection.First(&project, id)
	if result.Error != nil {
		return result.Error
	}

	project.Name = newProject.Name
	project.Description = newProject.Description
	saveResult := connection.Save(&project)

	if saveResult.Error != nil {
		return saveResult.Error
	}

	return nil
}

func Delete(id uint64) error {
	var project Project

	result := connection.First(&project, id)
	if result.Error != nil {
		return result.Error
	}

	deleteResult := connection.Delete(&project)
	if deleteResult.Error != nil {
		return deleteResult.Error
	}

	return nil
}
