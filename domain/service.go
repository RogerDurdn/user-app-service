package domain

import (
	"github.com/RogerDurdn/users/errors"
	"github.com/RogerDurdn/users/model"
	"github.com/RogerDurdn/users/pkg"
)

type Service interface {
	FindUserById(id int) (*model.User, error)
	FindUserByUserName(userName string) (*model.User, error)
	CreateOrUpdateUser(user *model.User) (*model.User, error)
	DeleteUserById(id int) error
	AuthUser(userName, pwd string) (*model.User, error)
}

type NormalSrv struct {
	storage pkg.Storage
}

func NewNormalSrv(storage pkg.Storage) *NormalSrv {
	return &NormalSrv{storage: storage}
}

func (ns *NormalSrv) FindUserById(id int) (*model.User, error) {
	user, err := ns.storage.FindUserById(id)
	if err != nil {
		return user, errors.NotFoundError("Not found user with id:" + err.Error())
	}
	return user, err
}

func (ns *NormalSrv) FindUserByUserName(userName string) (*model.User, error) {
	user, err := ns.storage.FindUserByUserName(userName)
	if err != nil {
		return user, errors.NotFoundError("Not found user with name:" + err.Error())
	}
	return user, err
}

func (ns *NormalSrv) CreateOrUpdateUser(user *model.User) (*model.User, error) {
	user, err := ns.storage.CreateOrUpdateUser(user)
	if err != nil {
		return user, errors.InternalServerError("Cannot create user:" + err.Error())
	}
	return user, nil
}

func (ns *NormalSrv) DeleteUserById(id int) error {
	if err := ns.storage.DeleteUserById(id); err == nil {
		return err
	}
	return errors.InternalServerError("Cannot delete user")
}

func (ns *NormalSrv) AuthUser(userName, pwd string) (*model.User, error) {
	if user, err := ns.FindUserByUserName(userName); err == nil {
		if user.Password == pwd {
			return user, nil
		}
	}
	return nil, errors.Unauthorized("nop")
}
