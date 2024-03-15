package port

import (
	port_shared "mc-stove/shared/port"
)

type StovePlantIRepository interface {
	port_shared.IRepositoryCRUD
}
