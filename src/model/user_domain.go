package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Ricardolv/mvc-api/src/config/rest_err"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{email, password, name, age}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	Create() *rest_err.RestErr
	Update(string) *rest_err.RestErr
	FindBy(string) (*UserDomain, *rest_err.RestErr)
	Delete(string) *rest_err.RestErr
}
