package permisions

import (
	"github.com/les-cours/user-api/types"
)

func CanRead(user *types.UserToken) bool {

	if !user.Read.LEARNING {
		return false
	}

	return true
}
