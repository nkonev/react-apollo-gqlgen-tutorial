package store

import (
	"context"
	"net/http"
	model "react-apollo-gqlgen-tutorial/backoffice/models"
)

// AuthQuery - queryResolver
// Возвращает доступную модель Auth. GET запрос
func (s *Store) AuthQuery(ctx context.Context) (*model.Auth, error) {
	auth := &model.Auth{}
	// ...
	return auth, nil
}

// AuthWebsocket - subscriptionResolver
// Метод вызывается в pkg/graph/auth.go. Осуществляет подключение по Websocket
func (s *Store) AuthWebsocket(ctx context.Context) (<-chan *model.Auth, error) {
	ch := make(chan *model.Auth)
	// ...
	return ch, nil
}

// AuthorizeForUsername - mutationResolver
// Метод осуществляющий авторизацию по username
func (s *Store) AuthorizeForUsername(ctx context.Context, username string) (*model.Auth, error) {
	auth := &model.Auth{}
	// ...
	return auth, nil
}

// AuthVerifyCode - mutationResolver
// Метод осуществляющий подтверждение кода из СМС
func (s *Store) AuthVerifyCode(ctx context.Context, code string) (*model.Auth, error) {
	auth := &model.Auth{}
	// ...
	return auth, nil
}

// AuthorizationHTTP
// Здесь будем работать с HTTP заголовками и Cookie
func (s *Store) AuthorizationHTTP(w http.ResponseWriter, r *http.Request) *http.Request {
	ctx := r.Context()
	// ...
	return r.WithContext(ctx)
}

// Cors
// Здесь будем работать с Cors
func (s *Store) Cors(w http.ResponseWriter, r *http.Request) *http.Request {
	ctx := r.Context()
	// ...
	return r.WithContext(ctx)
}
