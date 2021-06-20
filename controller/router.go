package controller

type Method string

const (
	REQUEST_METHOD_GET    Method = "GET"
	REQUEST_METHOD_PUT    Method = "PUT"
	REQUEST_METHOD_POST   Method = "POST"
	REQUEST_METHOD_DELETE Method = "DELETE"
)

type Route struct {
	Name        string
	Method      Method
	Path        string
	HandlerFunc func(Context) error
}

type Routes []Route

func (routes *Routes) ToArray() []Route {
	return ([]Route)(*routes)
}

func (routes *Routes) Append(other Routes) Routes {
	return (Routes)(append(routes.ToArray(), other.ToArray()...))
}
