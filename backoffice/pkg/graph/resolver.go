package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"react-apollo-gqlgen-tutorial/backoffice/graph/generated"
	"react-apollo-gqlgen-tutorial/backoffice/pkg/store"
)

type Resolver struct{
	
	// Подключим стор
	store *store.Store
}

// Создадим функцию New
func NewServer(opt Options) *handler.Server {
	
	return handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &Resolver{
					store: opt.Store,
				},
			},
		),
	)
}

type Options struct {
	Store *store.Store
}