package user

import (
	"time"

	"github.com/mkaiho/go-deploy-sample/domain/model"
)

type User struct {
	ID        model.Id
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserJson struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser() *User {
	id := model.NewId()
	return &User{
		ID: *id,
	}
}

func (user *User) ToUserJson() *UserJson {
	return &UserJson{
		ID:        user.ID.Value(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
