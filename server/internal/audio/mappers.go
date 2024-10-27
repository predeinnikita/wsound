package audio

func FromDalToEntity(dal AudioDal) AudioEntity {
	return AudioEntity{
		ID:        dal.ID,
		Name:      dal.Name,
		StorageID: dal.StorageID,
		ProjectID: dal.ProjectID,
	}
}

func FromEntityToDal(entity AudioEntity) AudioDal {
	return AudioDal{
		ID:        entity.ID,
		Name:      entity.Name,
		StorageID: entity.StorageID,
		ProjectID: entity.ProjectID,
	}
}
