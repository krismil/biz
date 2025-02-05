package utils

import "context"

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId := ctx.Value(SessionUserId)
	if userId == nil {
		return 0
	}
	return userId.(int32)
}
