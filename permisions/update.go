package permisions

import (
	"github.com/les-cours/user-api/types"
)

func CanUpdate(user *types.UserToken) bool {

	if !user.Update.LEARNING {
		return false
	}

	return true
}
