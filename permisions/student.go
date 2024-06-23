package permisions

import (
	"context"
	"github.com/les-cours/user-api/types"
	"log"
)

func Student(ctx context.Context) (*types.UserToken, error) {

	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return nil, ErrAuth
	}

	log.Println("suerTyp:", user.UserType)
	if user.UserType != "student" {
		return nil, ErrPermissionDenied
	}

	return user, nil
}
