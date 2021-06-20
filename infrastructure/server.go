package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/mkaiho/go-deploy-sample/controller"
	"github.com/mkaiho/go-deploy-sample/repository"
	"github.com/mkaiho/go-deploy-sample/service"
	"github.com/mkaiho/go-deploy-sample/types/server"
)

func NewServer() server.Server {
	return &echoServer{}
}

type echoServer struct {
	e  *echo.Echo
	ds *repository.DatasourceHandler
}

func (server *echoServer) Init() error {
	if server.e == nil {
		server.e = echo.New()
	}

	if ds, err := NewDatasource(); err != nil {
		return err
	} else if server.ds == nil {
		server.ds = &ds
	}

	routes := allRoutes(*server.ds)
	for _, route := range routes.ToArray() {
		var returnedRoute *echo.Route
		handlerFunc := route.HandlerFunc
		switch route.Method {
		case controller.REQUEST_METHOD_GET:
			returnedRoute = server.e.GET(route.Path, func(c echo.Context) error {
				return handlerFunc(c)
			})
		case controller.REQUEST_METHOD_PUT:
			returnedRoute = server.e.PUT(route.Path, func(c echo.Context) error {
				return handlerFunc(c)
			})
		case controller.REQUEST_METHOD_POST:
			returnedRoute = server.e.POST(route.Path, func(c echo.Context) error {
				return handlerFunc(c)
			})
		case controller.REQUEST_METHOD_DELETE:
			returnedRoute = server.e.DELETE(route.Path, func(c echo.Context) error {
				return handlerFunc(c)
			})
		}
		if route.Name != "" {
			returnedRoute.Name = route.Name
		}
	}

	return nil
}

func (server *echoServer) Run() error {
	if err := server.Init(); err != nil {
		return err
	}

	server.e.Logger.Fatal(server.e.Start(":3000"))

	return nil
}

func allRoutes(ds repository.DatasourceHandler) controller.Routes {
	routes := controller.Routes{}

	routes = routes.Append(*controller.UserRoutes(controller.NewUserController(service.NewUserService(repository.NewUserRepository(ds)))))
	return routes
}
