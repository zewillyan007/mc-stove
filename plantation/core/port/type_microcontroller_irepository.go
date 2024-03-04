package port

import (
	port_shared "mc-stove/shared/port"
)

type TypeMicrocontrollerIRepository interface {
	port_shared.IRepositoryCRUD
}
