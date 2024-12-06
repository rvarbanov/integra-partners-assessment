package api

import (
	"backend/internal/config"
	"backend/internal/controller"
	"backend/internal/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	ctrl controller.ControllerInterface
}

func NewAPI(ctrl controller.ControllerInterface) *API {
	return &API{ctrl: ctrl}
}

func (api *API) Start(cfg config.API) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/api/status", api.getStatus)

	e.POST("/api/user", api.createUser)
	e.GET("/api/user/:id", api.getUser)
	e.PUT("/api/user/:id", api.updateUser, validateUserBody)
	e.DELETE("/api/user/:id", api.deleteUser)

	e.GET("/api/users", api.getUsers)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

func (api *API) getStatus(c echo.Context) error {
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	return c.JSON(http.StatusOK, resp)
}

func (api *API) createUser(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	ID, err := api.ctrl.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: ID})
}

func (api *API) getUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	user, err := api.ctrl.GetUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: user})
}

func (api *API) updateUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	user := model.User{}
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	user, err = api.ctrl.UpdateUser(ID, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: user})
}

func (api *API) deleteUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	err = api.ctrl.DeleteUser(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func (api *API) getUsers(c echo.Context) error {
	users, err := api.ctrl.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: users})
}

func getParamID(c echo.Context) (int, error) {
	id := c.Param("id")
	if id == "" {
		return 0, fmt.Errorf("id is required")
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	if ID <= 0 {
		return 0, fmt.Errorf("id must be a positive integer")
	}

	return ID, nil
}

func validateUserBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.Response{Error: err.Error()})
		}

		if !user.IsValid() {
			return c.JSON(http.StatusBadRequest, model.Response{Error: "invalid user data"})
		}

		return next(c)
	}

}
