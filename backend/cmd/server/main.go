package main

import (
	"backend/internal/api"
	"backend/internal/config"
	"backend/internal/controller"
	"backend/internal/dao"
	"backend/internal/db"
)

func main() {
	cfg := config.Load()

	db := db.NewDB(cfg.Database)

	dao := dao.NewDao(db)

	controller := controller.NewController(dao)

	api := api.NewAPI(controller)

	api.Start(cfg.API)

}
