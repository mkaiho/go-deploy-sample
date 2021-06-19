package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/mkaiho/go-deploy-sample/controller"
	"github.com/mkaiho/go-deploy-sample/repository"
	"github.com/mkaiho/go-deploy-sample/service"
)

type handler struct{}

func NewHandler() controller.Handler {
	return &handler{}
}

func (h *handler) Run() {
	e := echo.New()
	ds, _ := NewDatasource()
	userController := controller.NewUserController(service.NewUserService(repository.NewUserRepository(ds)))
	e.GET("/users/:id", func(c echo.Context) error {
		userController.Get(c)
		return nil
	})
	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}
