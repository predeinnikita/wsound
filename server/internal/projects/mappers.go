package projects

func FromDalToEntity(dal ProjectDal) ProjectEntity {
	return ProjectEntity{
		ID:          dal.ID,
		Name:        dal.Name,
		Description: dal.Description,
		CreatedAt:   dal.CreatedAt,
	}
}

func FromEntityToDal(entity ProjectEntity) ProjectDal {
	return ProjectDal{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		CreatedAt:   entity.CreatedAt,
	}
}
