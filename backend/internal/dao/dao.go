package dao

import (
	"backend/internal/db"
	"backend/internal/model"
)

type DaoInterface interface {
	GetUsers() ([]model.User, error)
	GetUser(ID int) (model.User, error)
	CreateUser(user model.User) (int, error)
	UpdateUser(ID int, user model.User) (model.User, error)
	DeleteUser(ID int) error
	GetUserByUsername(username string) (model.User, error)
}

type Dao struct {
	DB db.DBInterface
}

func NewDao(db db.DBInterface) *Dao {
	return &Dao{DB: db}
}

func (d *Dao) GetUsers() ([]model.User, error) {
	return d.DB.GetUsers()
}

func (d *Dao) GetUser(ID int) (model.User, error) {
	return d.DB.GetUser(ID)
}

func (d *Dao) CreateUser(user model.User) (int, error) {
	return d.DB.CreateUser(user)
}

func (d *Dao) UpdateUser(ID int, user model.User) (model.User, error) {
	return d.DB.UpdateUser(ID, user)
}

func (d *Dao) DeleteUser(ID int) error {
	return d.DB.DeleteUser(ID)
}

func (d *Dao) GetUserByUsername(username string) (model.User, error) {
	return d.DB.GetUserByUsername(username)
}
