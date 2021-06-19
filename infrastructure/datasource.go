package infrastructure

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mkaiho/go-deploy-sample/repository"
)

func NewDatasource() (repository.DatasourceHandler, error) {
	return newMysqlDatasource("devuser", "devdev", "mysqldb", 3306, "devdb")
}

func newMysqlDatasource(user string, password string, host string, port int, dbname string) (repository.DatasourceHandler, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, dbname))
	if err != nil {
		return nil, err
	}
	return &mysqlDatasource{db}, nil
}

type mysqlDatasource struct {
	db *sql.DB
}

func (datasource *mysqlDatasource) Execute(query string, args ...interface{}) error {
	stmt, err := datasource.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (datasource *mysqlDatasource) Query(query string, args ...interface{}) (repository.Rows, error) {
	stmt, err := datasource.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ctypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	types := make([]reflect.Type, len(ctypes))
	for i, c := range ctypes {
		switch c.ScanType().String() {
		case "sql.RawBytes":
			types[i] = reflect.TypeOf("")
		case "sql.NullTime":
			types[i] = reflect.TypeOf(time.Now())
		default:
			types[i] = c.ScanType()
		}
	}

	values := make([][]interface{}, 0)
	for rows.Next() {
		row := make([]interface{}, len(types))
		for i, t := range types {
			row[i] = reflect.New(t).Interface()
		}

		if err := rows.Scan(row...); err != nil {
			return nil, err
		}
		for i, v := range row {
			row[i] = reflect.ValueOf(v).Elem().Interface()
		}
		values = append(values, row)
	}
	return values, nil
}

func (datasource *mysqlDatasource) Close() error {
	return datasource.db.Close()
}
