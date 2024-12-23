package model

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
	SetID(string)
	GetID() string
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{email: email, password: password, name: name, age: age}
}