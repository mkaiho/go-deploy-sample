package controller

type Context interface {
	Bind(i interface{}) error
	FormValue(name string) string
	Param(name string) string
	JSON(code int, i interface{}) error
}
