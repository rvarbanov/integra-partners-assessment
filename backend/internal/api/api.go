package api

import (
	"backend/internal/config"
	"backend/internal/controller"
	"backend/internal/model"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "backend/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/api/v1/status", api.getStatus)

	e.POST("/api/v1/user", api.createUser)
	e.GET("/api/v1/user/:id", api.getUser)
	e.PUT("/api/v1/user/:id", api.updateUser)
	e.DELETE("/api/v1/user/:id", api.deleteUser)

	e.GET("/api/v1/users", api.getUsers)

	e.GET("/api/v1/docs/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

// @Summary Get the status of the API
// @Description Returns the status of the API
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=string} "OK"
// @Router /api/v1/status [get]
func (api *API) getStatus(c echo.Context) error {
	resp := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	return c.JSON(http.StatusOK, resp)
}

// @Summary Create a new user
// @Description Creates a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 201 {object} model.Response{data=int} "User created successfully, returns user ID"
// @Failure 400 {object} model.Response{error=string} "Invalid request body"
// @Failure 409 {object} model.Response{error=string} "Username already exists"
// @Failure 500 {object} model.Response{error=string} "Internal server error"
// @Router /api/v1/user [post]
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
		if strings.Contains(err.Error(), "username already exists") {
			return c.JSON(http.StatusConflict, model.Response{
				Error: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	resp := model.Response{
		Data: model.User{
			ID: ID,
		},
	}
	return c.JSON(http.StatusCreated, resp)
}

// @Summary Get user by ID
// @Description Retrieves a user by their ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.Response{data=model.User} "User found"
// @Failure 400 {object} model.Response{error=string} "Invalid user ID"
// @Failure 404 {object} model.Response{error=string} "User not found"
// @Failure 500 {object} model.Response{error=string} "Internal server error"
// @Router /api/v1/user/{id} [get]
func (api *API) getUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: err.Error(),
		})
	}

	user, err := api.ctrl.GetUser(ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, model.Response{
				Error: err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: user})
}

// @Summary Update a user by ID
// @Description Updates a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "User object"
// @Success 200 {object} model.Response{data=model.User} "User updated successfully"
// @Failure 400 {object} model.Response{error=string} "Bad Request"
// @Failure 404 {object} model.Response{error=string} "User not found"
// @Failure 500 {object} model.Response{error=string} "Internal Server Error"
// @Router /api/v1/user/{id} [put]
func (api *API) updateUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: fmt.Errorf("invalid user ID: %w", err).Error(),
		})
	}

	user := model.User{}
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: fmt.Errorf("invalid user data: %w", err).Error(),
		})
	}

	user, err = api.ctrl.UpdateUser(ID, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: fmt.Errorf("failed to update user: %w", err).Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{Data: user})
}

// @Summary Delete a user by ID
// @Description Deletes a user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "User deleted successfully"
// @Failure 400 {object} model.Response{error=string} "Bad Request"
// @Failure 404 {object} model.Response{error=string} "User not found"
// @Failure 500 {object} model.Response{error=string} "Internal Server Error"
// @Router /api/v1/user/{id} [delete]
func (api *API) deleteUser(c echo.Context) error {
	ID, err := getParamID(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Error: fmt.Errorf("invalid user ID: %w", err).Error(),
		})
	}

	err = api.ctrl.DeleteUser(ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, model.Response{
				Error: fmt.Errorf("user not found: %w", err).Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: fmt.Errorf("failed to delete user: %w", err).Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// @Summary Get all users
// @Description Returns all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=[]model.User} "Users found"
// @Failure 500 {object} model.Response{error=string} "Internal Server Error"
// @Router /api/v1/users [get]
func (api *API) getUsers(c echo.Context) error {
	users, err := api.ctrl.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Error: fmt.Errorf("failed to get users: %w", err).Error(),
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

	return ID, err
}

/*
func validateUserBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.User{}
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.Response{Error: fmt.Errorf("invalid user data: %w", err.Error())})
		}

		if !user.IsValid() {
			return c.JSON(http.StatusBadRequest, model.Response{Error: fmt.Errorf("invalid user data: %w", err.Error())})
		}

		return next(c)
	}

}
*/
