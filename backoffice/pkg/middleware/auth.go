package middleware

import (
	"net/http"
	"react-apollo-gqlgen-tutorial/backoffice/pkg/store"
)


func AuthMiddleware(store *store.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Подключим обработчик авторизации из Store
			// Получим в ответе контексте
			r = store.AuthorizationHTTP(w, r)


			// Вернем контекст
			next.ServeHTTP(w, r)
		})
	}
}

