package port

import (
	port_shared "mc-stove/shared/port"
)

type CurrentParamsIRepository interface {
	port_shared.IRepositoryCRUD
}
