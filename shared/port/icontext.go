package port

import "mc-stove/shared/types"

type IManagerContext interface {
	GetUser() *types.UserContext
}
