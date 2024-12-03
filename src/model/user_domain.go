package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{email, password, name, age}
}

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) GetName() string {

	return ud.name
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))

	ud.password = hex.EncodeToString(hash.Sum(nil))
}
