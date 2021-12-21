package store

import (
	"context"
	model "react-apollo-gqlgen-tutorial/backoffice/models"
)

// UserQuery - queryResolver
// Возвращает доступную модель User. GET запрос
func (s *Store) UserQuery(ctx context.Context) (*model.User, error) {
	user := &model.User{}
	// ...
	return user, nil
}
