package repository

import (
	"fmt"
	"time"

	model "github.com/mkaiho/go-deploy-sample/domain/model"
	user "github.com/mkaiho/go-deploy-sample/domain/model/user"
)

const (
	allColumnNameClause = "id, name, created_at, updated_at"
)

type UserRepository interface {
	FindAll() (users []*user.User, err error)
	Find(id string) (user *user.User, err error)
	Create(users *user.User) error
	Update(users *user.User) error
}

type userRepository struct {
	datasource DatasourceHandler
}

func NewUserRepository(datasource DatasourceHandler) UserRepository {
	return &userRepository{datasource: datasource}
}

func (repository *userRepository) FindAll() ([]*user.User, error) {
	query := fmt.Sprintf("SELECT %v FROM users", allColumnNameClause)
	rows, err := repository.datasource.Query(query)
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)
	for _, values := range rows {
		id, err := model.ValueOf(values[0].(string))
		if err != nil {
			return nil, err
		}

		users = append(users, &user.User{
			ID:        *id,
			Name:      values[1].(string),
			CreatedAt: values[2].(time.Time),
			UpdatedAt: values[3].(time.Time),
		})
	}
	return users, nil
}

func (repository *userRepository) Find(id string) (*user.User, error) {
	rows, err := repository.datasource.Query(fmt.Sprintf("SELECT %v FROM users WHERE id = ?", allColumnNameClause), id)
	if err != nil {
		return nil, err
	}
	if len(rows) > 1 {
		return nil, fmt.Errorf("too many rows was fetched")
	}
	if len(rows) == 0 {
		return nil, nil
	}

	userId, err := model.ValueOf(rows[0][0].(string))
	if err != nil {
		return nil, err
	}

	user := &user.User{
		ID:        *userId,
		Name:      rows[0][1].(string),
		CreatedAt: rows[0][2].(time.Time),
		UpdatedAt: rows[0][3].(time.Time),
	}

	return user, nil
}

func (repository *userRepository) Create(user *user.User) error {
	query := fmt.Sprintf("INSERT INTO users (id, name) VALUES (?, ?)")

	if err := repository.datasource.Execute(query, user.ID.Value(), user.Name); err != nil {
		return err
	}

	return nil
}

func (repository *userRepository) Update(user *user.User) error {
	query := "UPDATE users SET name = ? WHERE id = ?"

	if err := repository.datasource.Execute(query, user.Name, user.ID.Value()); err != nil {
		return err
	}

	return nil
}
