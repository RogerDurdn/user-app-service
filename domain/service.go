package domain

import (
	"github.com/RogerDurdn/users/model"
	"github.com/RogerDurdn/users/pkg"
)

type Service interface {
	FindUserById(id int) (*model.User, *model.ErrorWrap)
	FindUserByName(name string) (*model.User, *model.ErrorWrap)
	DeleteUserById(id int) (bool, *model.ErrorWrap)
	AuthUser(userName, pwd string) (bool, *model.ErrorWrap)
	CreateOrUpdateUser(user *model.User) (*model.User, *model.ErrorWrap)
}

type NormalSrv struct {
	storage pkg.Storage
}

func NewNormalSrv(storage pkg.Storage) *NormalSrv {
	return &NormalSrv{storage: storage}
}

func (ns *NormalSrv) FindUserById(id int) (*model.User, *model.ErrorWrap) {
	return &model.User{Id: 123, Name: "roger", Password: "secure"}, nil
}

func (ns *NormalSrv) FindUserByName(name string) (*model.User, *model.ErrorWrap) {
	return nil, nil
}

func (ns *NormalSrv) DeleteUserById(id int) (bool, *model.ErrorWrap) {
	return false, nil
}

func (ns *NormalSrv) AuthUser(userName, pwd string) (bool, *model.ErrorWrap) {
	return false, nil
}

func (ns *NormalSrv) CreateOrUpdateUser(user *model.User) (*model.User, *model.ErrorWrap) {
	return nil, nil
}
