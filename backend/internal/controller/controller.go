package controller

import (
	"backend/internal/dao"
	"backend/internal/model"
	"fmt"
)

type ControllerInterface interface {
	CreateUser(user model.User) (int, error)
	GetUser(ID int) (model.User, error)
	UpdateUser(ID int, user model.User) (model.User, error)
	DeleteUser(ID int) error
	GetUsers() ([]model.User, error)
}

type Controller struct {
	dao dao.DaoInterface
}

func NewController(dao dao.DaoInterface) *Controller {
	return &Controller{dao: dao}
}

func (ctrl *Controller) CreateUser(user model.User) (int, error) {
	// return an error if the username is already taken
	userWithMatchingUsername, err := ctrl.GetUserByUsername(user.Username)
	if err != nil {
		return 0, err
	}

	if userWithMatchingUsername.Username == user.Username {
		return 0, fmt.Errorf("username already exists")
	}

	return ctrl.dao.CreateUser(user)
}

func (ctrl *Controller) GetUser(ID int) (model.User, error) {
	return ctrl.dao.GetUser(ID)
}

func (ctrl *Controller) UpdateUser(ID int, user model.User) (model.User, error) {
	existingUser, err := ctrl.GetUser(ID)
	if err != nil {
		return model.User{}, err
	}

	if existingUser.Username == user.Username {
		return ctrl.dao.UpdateUser(ID, user)
	}

	userWithMatchingUsername, err := ctrl.GetUserByUsername(user.Username)
	if err != nil {
		return model.User{}, err
	}

	if userWithMatchingUsername.ID != 0 {
		return model.User{}, fmt.Errorf("username already exists")
	}

	return ctrl.dao.UpdateUser(ID, user)
}

func (ctrl *Controller) DeleteUser(ID int) error {
	return ctrl.dao.DeleteUser(ID)
}

func (ctrl *Controller) GetUsers() ([]model.User, error) {
	return ctrl.dao.GetUsers()
}

func (ctrl *Controller) GetUserByUsername(username string) (model.User, error) {
	return ctrl.dao.GetUserByUsername(username)
}
