package repository

type Rows [][]interface{}

type DatasourceHandler interface {
	Execute(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (Rows, error)
	Close() error
}
