package model

import (
	"context"
)

type Meta struct {
	Uid 		int
	Cid 		string
	Role 		string
	Reconnect 	bool
	Authorized 	bool
}

type key struct {
	name string
}

func (m *Meta) Value(ctx context.Context) *Meta {
	meta := ctx.Value(key{"meta"})
	if meta == nil {
		return m
	}
	return meta.(*Meta)
}

func (m *Meta) WithContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, key{"meta"}, m)
	return ctx
}