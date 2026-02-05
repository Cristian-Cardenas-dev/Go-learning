package port

import "api-gateway/pkg/domain"

type UserRepository interface {
	Insert(user *domain.User) (err error)
}
