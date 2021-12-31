package store

import (
	"context"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	model "react-apollo-gqlgen-tutorial/backoffice/models"
	"time"
)

// AuthQuery - queryResolver
// Возвращает доступную модель Auth. GET запрос
func (s *Store) AuthQuery(ctx context.Context) (*model.Auth, error) {
	// Получим мета из контекста
	meta := &model.Meta{}
	meta = meta.Value(ctx)

	auth 			:= &model.Auth{}
	auth.Reconnect 	= meta.Reconnect

	return auth, nil
}

// AuthWebsocket - subscriptionResolver
// Метод вызывается в pkg/graph/auth.go. Осуществляет подключение по Websocket
func (s *Store) AuthWebsocket(ctx context.Context) (<-chan *model.Auth, error) {
	// Получим мета из контекста
	meta := &model.Meta{}
	meta = meta.Value(ctx)

	// Создадим websocket id
	wsid := uuid.New().String()

	// Создадим канал
	ch := make(chan *model.Auth)

	// Логика по добавлению слушателя
	fmt.Printf("Connect CID: %v, WSID: %v\n", meta.Cid, wsid)

	// Логика по удалению слушателя
	go func() {
		<- ctx.Done()
		fmt.Printf("Disconnect CID: %v, WSID: %v\n", meta.Cid, wsid)
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch <- &model.Auth{
			Method: "method",
		}
	}()

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

	// Создадим мету
	meta := &model.Meta{}

	// Проверим наличие токена Cid
	cidCookie, err := r.Cookie("_cid")
	if err != nil {
		meta.Reconnect = true

		// Токена нет, создадим
		cid := uuid.New().String()
		cidCookie = &http.Cookie{
			Name: "_cid",
			Value: cid,
			HttpOnly: true,
			Secure: true,
		}
		http.SetCookie(w, cidCookie)
	}

	// Прочитаем Cid и запишем в контекст
	if cidCookie != nil {
		meta.Cid = cidCookie.Value
	}

	return r.WithContext(
		// Запишем мета в контекст
		meta.WithContext(ctx),
	)
}

// Cors
// Здесь будем работать с Cors
func (s *Store) Cors(w http.ResponseWriter, r *http.Request) *http.Request {
	ctx := r.Context()
	// ...
	return r.WithContext(ctx)
}
