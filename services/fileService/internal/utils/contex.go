package utils

import "context"

type CtxKey string

var (
	userTokenKey CtxKey = "userToken"
)

func WithUserToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, userTokenKey, token)
}

func GetUserToken(ctx context.Context) (token string, ok bool) {
	token, ok = ctx.Value(userTokenKey).(string)
	return
}
