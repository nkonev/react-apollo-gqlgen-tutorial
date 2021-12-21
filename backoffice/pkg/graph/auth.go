package graph

import (
	"context"
	model "react-apollo-gqlgen-tutorial/backoffice/models"
)


func (r *queryResolver) Auth(ctx context.Context) (*model.Auth, error) {
	return r.store.AuthQuery(ctx)
}

func (r *mutationResolver) AuthLogin(ctx context.Context, login string) (*model.Auth, error) {
	return r.store.AuthorizeForUsername(ctx, login)
}

func (r *mutationResolver) AuthVerifyCode(ctx context.Context, code string) (*model.Auth, error) {
	return r.store.AuthVerifyCode(ctx, code)
}

func (r *subscriptionResolver) AuthSubscription(ctx context.Context) (<-chan *model.Auth, error) {
	return r.store.AuthWebsocket(ctx)
}
