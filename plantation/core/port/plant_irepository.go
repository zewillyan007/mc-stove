package port

import (
	port_shared "mc-stove/shared/port"
)

type PlantIRepository interface {
	port_shared.IRepositoryCRUD
}
