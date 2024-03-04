package port

import (
	port_shared "mc-stove/shared/port"
)

type MicrocontrollerIRepository interface {
	port_shared.IRepositoryCRUD
}
