package permisions

import (
	"github.com/les-cours/user-api/types"
)

func CanDelete(user *types.UserToken) bool {

	if !user.Delete.LEARNING {
		return false
	}

	return true
}
