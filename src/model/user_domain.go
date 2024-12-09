package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
	GetJSONValue() (string, error)
	SetID(string)
}

func NewUserDomain(email, password, name string, age int8) *userDomain {
	return &userDomain{Email: email, Password: password, Name: name, Age: age}
}

type userDomain struct {
	ID       string
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"Name"`
	Age      int8   `json:"Age"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(b), nil
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id

}

func (ud *userDomain) GetName() string {

	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))

	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
