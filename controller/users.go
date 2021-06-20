package controller

import (
	"errors"
	"net/http"

	"github.com/mkaiho/go-deploy-sample/domain/model/user"
	"github.com/mkaiho/go-deploy-sample/service"
)

type UserController interface {
	Get(c Context) error
	Create(c Context) error
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func UserRoutes(controller UserController) *Routes {
	return &Routes{
		Route{
			Method:      REQUEST_METHOD_POST,
			Path:        "/users",
			HandlerFunc: controller.Create,
		},
		Route{
			Method:      REQUEST_METHOD_GET,
			Path:        "/users/:id",
			HandlerFunc: controller.Get,
		},
	}
}

func (controller *userController) Get(c Context) error {
	id := c.Param("id")
	u, err := controller.userService.Find(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{})
	}
	return c.JSON(http.StatusOK, u.ToUserJson())
}

func (controller *userController) Create(c Context) error {
	json := map[string]interface{}{}
	c.Bind(&json)

	if json["name"] == nil || json["name"] == "" {
		return c.JSON(http.StatusBadRequest, NewError(errors.New("Error")))
	}

	user := user.NewUser()
	user.Name = json["name"].(string)
	err := controller.userService.Create(*user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err))
	}

	user, err = controller.userService.Find(user.ID.Value())
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err))
	}

	return c.JSON(http.StatusOK, user.ToUserJson())
}
