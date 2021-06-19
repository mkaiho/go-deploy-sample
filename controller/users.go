package controller

import (
	"errors"
	"net/http"

	"github.com/mkaiho/go-deploy-sample/domain/model/user"
	"github.com/mkaiho/go-deploy-sample/service"
)

type UserController interface {
	Get(c Context)
	Create(c Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (controller *userController) Get(c Context) {
	id := c.Param("id")
	u, err := controller.userService.Find(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(err))
		return
	}
	if u == nil {
		c.JSON(http.StatusOK, map[string]interface{}{})
		return
	}
	c.JSON(http.StatusOK, u.ToUserJson())
}

func (controller *userController) Create(c Context) {
	json := map[string]interface{}{}
	c.Bind(&json)

	if json["name"] == nil || json["name"] == "" {
		c.JSON(http.StatusBadRequest, NewError(errors.New("Error")))
		return
	}

	user := user.NewUser()
	user.Name = json["name"].(string)
	err := controller.userService.Create(*user)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
	}

	user, err = controller.userService.Find(user.ID.Value())
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(err))
	}

	c.JSON(http.StatusOK, user.ToUserJson())
}
