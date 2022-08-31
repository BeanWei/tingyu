package shared

import "context"

type ctxKeyType string

// CtxSvcKey 上下文变量存储键名
const CtxSvcKey ctxKeyType = "CtxSvcKey"

// Ctx 请求上下文结构
type Ctx struct {
	User *CtxUser // 上下文用户信息
}

// CtxUser 上下文用户信息
type CtxUser struct {
	ID      int64 `json:"id"`
	IsAdmin bool  `json:"is_admin"`
}

func GetCtxUser(ctx context.Context) *CtxUser {
	if localCtx, ok := ctx.Value(CtxSvcKey).(*Ctx); ok {
		return localCtx.User
	}
	return nil
}
