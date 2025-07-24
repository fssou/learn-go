package lake

import (
	"github.com/fssou/learn-go/internal/domain/entity"
	"github.com/fssou/learn-go/internal/domain/port/out"
)

type Service interface {
	Save(lakeEntity *entity.LakeEntity) error
	RetrieveById(string) *entity.LakeEntity
}

type ServiceImpl struct {
	repository out.Repository[string, entity.LakeEntity]
}

func NewService(repository out.Repository[string, entity.LakeEntity]) *ServiceImpl {
	return &ServiceImpl{repository }
}

func (svc *ServiceImpl) Save(lakeEntity *entity.LakeEntity) error {
	return svc.repository.Save(lakeEntity)
}

func (svc *ServiceImpl) RetrieveById(id string) (*entity.LakeEntity, error) {
	return svc.repository.GetById(id)
}
