package services

import "demo/repositories"

type IUserService interface {
	Register() error
}

type UserService struct {
	IUserRepo repositories.IUserRepo
}

func (_self UserService) Register() error {
	return nil
}