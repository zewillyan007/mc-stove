package service

import (
	"mc-stove/plantation/core/domain/dto"
	"mc-stove/plantation/core/port"
	"mc-stove/plantation/core/usecase"
	"mc-stove/shared/grid"
	port_shared "mc-stove/shared/port"
	"mc-stove/shared/types"
	"time"
)

type TypeMicrocontrollerService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.TypeMicrocontrollerIRepository
	ucGet      *usecase.TypeMicrocontrollerUseCaseGet
	ucSave     *usecase.TypeMicrocontrollerUseCaseSave
	// ucGrid     *usecase.TypeMicrocontrollerUseCaseGrid
	ucGetAll *usecase.TypeMicrocontrollerUseCaseGetAll
	ucRemove *usecase.TypeMicrocontrollerUseCaseRemove
}

func NewTypeMicrocontrollerService(provider port_shared.IResourceProvider) *TypeMicrocontrollerService {
	repo := types.GetConstructor((*port.TypeMicrocontrollerIRepository)(nil))(provider.GetDB()).(port.TypeMicrocontrollerIRepository)
	repo.SetContext(provider.Context())

	return &TypeMicrocontrollerService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewTypeMicrocontrollerUseCaseGet(repo),
		ucSave:     usecase.NewTypeMicrocontrollerUseCaseSave(repo),
		// ucGrid:     usecase.NewTypeMicrocontrollerUseCaseGrid(repo),
		ucGetAll: usecase.NewTypeMicrocontrollerUseCaseGetAll(repo),
		ucRemove: usecase.NewTypeMicrocontrollerUseCaseRemove(repo),
	}
}

func (o *TypeMicrocontrollerService) WithTransaction(transaction port_shared.ITransaction) *TypeMicrocontrollerService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *TypeMicrocontrollerService) Get(dtoIn *dto.TypeMicrocontrollerDtoIn) (*dto.TypeMicrocontrollerDtoOut, error) {

	TypeMicrocontroller, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewTypeMicrocontrollerDtoOut()

	dtoOut.Id = TypeMicrocontroller.Id
	dtoOut.Name = TypeMicrocontroller.Name
	dtoOut.Mnemonic = TypeMicrocontroller.Mnemonic
	dtoOut.Hint = TypeMicrocontroller.Hint

	if TypeMicrocontroller.CreationDateTime != nil {
		dtoOut.CreationDateTime = TypeMicrocontroller.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if TypeMicrocontroller.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = TypeMicrocontroller.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *TypeMicrocontrollerService) GetAll(conditions ...interface{}) []*dto.TypeMicrocontrollerDtoOut {

	var arrayTypeMicrocontrollerDto []*dto.TypeMicrocontrollerDtoOut

	arrayTypeMicrocontroller := o.ucGetAll.Execute(conditions...)

	for _, TypeMicrocontroller := range arrayTypeMicrocontroller {

		dtoOut := dto.NewTypeMicrocontrollerDtoOut()

		dtoOut.Id = TypeMicrocontroller.Id
		dtoOut.Name = TypeMicrocontroller.Name
		dtoOut.Mnemonic = TypeMicrocontroller.Mnemonic
		dtoOut.Hint = TypeMicrocontroller.Hint

		if TypeMicrocontroller.CreationDateTime != nil {
			dtoOut.CreationDateTime = TypeMicrocontroller.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if TypeMicrocontroller.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = TypeMicrocontroller.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayTypeMicrocontrollerDto = append(arrayTypeMicrocontrollerDto, dtoOut)
	}
	return arrayTypeMicrocontrollerDto
}

func (o *TypeMicrocontrollerService) Save(dtoIn *dto.TypeMicrocontrollerDtoIn) error {

	var err error
	TypeMicrocontroller := FactoryTypeMicrocontroller()

	if dtoIn.Id > 0 {
		TypeMicrocontroller.Id = dtoIn.Id
	}

	TypeMicrocontroller.Name = dtoIn.Name
	TypeMicrocontroller.Mnemonic = dtoIn.Mnemonic
	TypeMicrocontroller.Hint = dtoIn.Hint

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if TypeMicrocontroller.Id == 0 {
			TypeMicrocontroller.CreationDateTime = &now
		} else {
			TypeMicrocontrollerCurrent, _ := o.ucGet.Execute(TypeMicrocontroller.Id)
			TypeMicrocontroller.CreationDateTime = TypeMicrocontrollerCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		TypeMicrocontroller.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		TypeMicrocontroller.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		TypeMicrocontroller.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(TypeMicrocontroller)
	if err != nil {
		return err
	}
	return nil
}

func (o *TypeMicrocontrollerService) Remove(dtoIn *dto.TypeMicrocontrollerDtoIn) error {

	TypeMicrocontroller := FactoryTypeMicrocontroller()
	if dtoIn.Id > 0 {
		TypeMicrocontroller.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(TypeMicrocontroller)
	if err != nil {
		return err
	}

	return nil
}

func (o *TypeMicrocontrollerService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	// dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
