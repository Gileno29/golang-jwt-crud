package converter

import (
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/repository/entity"
	"github.com/Gileno29/golang-jwt-crud/src/model"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}

}
