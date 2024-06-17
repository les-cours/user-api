package permisions

import (
	"context"
	"github.com/les-cours/user-api/types"
)

func User(ctx context.Context) (*types.UserToken, error) {

	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return nil, ErrAuth
	}

	return user, nil
}
