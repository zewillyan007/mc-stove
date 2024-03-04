package service

import (
	"mc-stove/plantation/core/domain/dto"
	"mc-stove/plantation/core/port"
	"mc-stove/plantation/core/usecase"
	port_shared "mc-stove/shared/port"
	"mc-stove/shared/types"
	"time"
)

type MicrocontrollerService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.MicrocontrollerIRepository
	ucGet      *usecase.MicrocontrollerUseCaseGet
	ucSave     *usecase.MicrocontrollerUseCaseSave
	// ucGrid     *usecase.MicrocontrollerUseCaseGrid
	ucGetAll *usecase.MicrocontrollerUseCaseGetAll
	ucRemove *usecase.MicrocontrollerUseCaseRemove
}

func NewMicrocontrollerService(provider port_shared.IResourceProvider) *MicrocontrollerService {
	repo := types.GetConstructor((*port.MicrocontrollerIRepository)(nil))(provider.GetDB()).(port.MicrocontrollerIRepository)
	repo.SetContext(provider.Context())

	return &MicrocontrollerService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewMicrocontrollerUseCaseGet(repo),
		ucSave:     usecase.NewMicrocontrollerUseCaseSave(repo),
		// ucGrid:     usecase.NewMicrocontrollerUseCaseGrid(repo),
		ucGetAll: usecase.NewMicrocontrollerUseCaseGetAll(repo),
		ucRemove: usecase.NewMicrocontrollerUseCaseRemove(repo),
	}
}

func (o *MicrocontrollerService) WithTransaction(transaction port_shared.ITransaction) *MicrocontrollerService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *MicrocontrollerService) Get(dtoIn *dto.MicrocontrollerDtoIn) (*dto.MicrocontrollerDtoOut, error) {

	Microcontroller, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewMicrocontrollerDtoOut()

	dtoOut.Id = Microcontroller.Id
	dtoOut.Type = Microcontroller.Type
	dtoOut.SerialNumber = Microcontroller.SerialNumber

	if Microcontroller.CreationDateTime != nil {
		dtoOut.CreationDateTime = Microcontroller.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if Microcontroller.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Microcontroller.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *MicrocontrollerService) GetAll(conditions ...interface{}) []*dto.MicrocontrollerDtoOut {

	var arrayMicrocontrollerDto []*dto.MicrocontrollerDtoOut

	arrayMicrocontroller := o.ucGetAll.Execute(conditions...)

	for _, Microcontroller := range arrayMicrocontroller {

		dtoOut := dto.NewMicrocontrollerDtoOut()

		dtoOut.Id = Microcontroller.Id
		dtoOut.Type = Microcontroller.Type
		dtoOut.SerialNumber = Microcontroller.SerialNumber

		if Microcontroller.CreationDateTime != nil {
			dtoOut.CreationDateTime = Microcontroller.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if Microcontroller.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Microcontroller.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayMicrocontrollerDto = append(arrayMicrocontrollerDto, dtoOut)
	}
	return arrayMicrocontrollerDto
}

func (o *MicrocontrollerService) Save(dtoIn *dto.MicrocontrollerDtoIn) error {

	var err error
	Microcontroller := FactoryMicrocontroller()

	if dtoIn.Id > 0 {
		Microcontroller.Id = dtoIn.Id
	}

	Microcontroller.Type = dtoIn.Type
	Microcontroller.SerialNumber = dtoIn.SerialNumber

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Microcontroller.Id == 0 {
			Microcontroller.CreationDateTime = &now
		} else {
			MicrocontrollerCurrent, _ := o.ucGet.Execute(Microcontroller.Id)
			Microcontroller.CreationDateTime = MicrocontrollerCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Microcontroller.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Microcontroller.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Microcontroller.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(Microcontroller)
	if err != nil {
		return err
	}
	return nil
}

func (o *MicrocontrollerService) Remove(dtoIn *dto.MicrocontrollerDtoIn) error {

	Microcontroller := FactoryMicrocontroller()
	if dtoIn.Id > 0 {
		Microcontroller.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(Microcontroller)
	if err != nil {
		return err
	}

	return nil
}

// func (o *MicrocontrollerService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
// 	var dataGrid map[string]interface{}
// 	var err error

// 	// if o._cache_ != nil {
// 	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
// 	// } else {
// 	dataGrid, err = o.ucGrid.Execute(GridConfig)
// 	// }

// 	return dataGrid, err
// }
