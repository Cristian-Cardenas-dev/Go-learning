package application

import (
	"api-gateway/cmd/api/core"
	"api-gateway/pkg/domain"
	"api-gateway/pkg/port"
	"fmt"
)

type UserCreator struct {
	Repo port.UserRepository
}

func NewUserCreator(repo port.UserRepository) *UserCreator {
	return &UserCreator{
		Repo: repo,
	}
}

func (uc *UserCreator) Execute(userParams core.UserCreateParams) string {
	user := &domain.User{
		Id:   userParams.Id,
		Name: userParams.Name,
		Age:  userParams.Age,
	}
	err := uc.Repo.Insert(user)
	fmt.Println(err)
	return "operating"
}
