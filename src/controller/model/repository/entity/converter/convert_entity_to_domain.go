package converter

import (
	"github.com/Gileno29/golang-jwt-crud/src/controller/model/repository/entity"
	"github.com/Gileno29/golang-jwt-crud/src/model"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age)

	domain.SetID(entity.ID.Hex())

	return domain
}
