package infrastructure

import (
	"context"
	"time"

	"flamingo.me/flamingo-app/src/user/domain"
)

type User struct {
	ID           uint
	Name         string
	EmailAddress *string
	ApprovedAt   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// UserServiceImpl default implementation
type UserServiceImpl struct{}

// UserByID mocked user getter
func (us *UserServiceImpl) UserByID(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{
		Name:      "User " + id,
		Nicknames: []string{"nick", id},
		Attributes: map[string]interface{}{
			"movie": "starwars",
		},
	}, nil
}
