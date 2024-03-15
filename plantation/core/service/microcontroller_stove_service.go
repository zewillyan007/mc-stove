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

type MicrocontrollerStoveService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.MicrocontrollerStoveIRepository
	ucGet      *usecase.MicrocontrollerStoveUseCaseGet
	ucSave     *usecase.MicrocontrollerStoveUseCaseSave
	// ucGrid     *usecase.MicrocontrollerStoveUseCaseGrid
	ucGetAll *usecase.MicrocontrollerStoveUseCaseGetAll
	ucRemove *usecase.MicrocontrollerStoveUseCaseRemove
}

func NewMicrocontrollerStoveService(provider port_shared.IResourceProvider) *MicrocontrollerStoveService {
	repo := types.GetConstructor((*port.MicrocontrollerStoveIRepository)(nil))(provider.GetDB()).(port.MicrocontrollerStoveIRepository)
	repo.SetContext(provider.Context())

	return &MicrocontrollerStoveService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewMicrocontrollerStoveUseCaseGet(repo),
		ucSave:     usecase.NewMicrocontrollerStoveUseCaseSave(repo),
		// ucGrid:     usecase.NewMicrocontrollerStoveUseCaseGrid(repo),
		ucGetAll: usecase.NewMicrocontrollerStoveUseCaseGetAll(repo),
		ucRemove: usecase.NewMicrocontrollerStoveUseCaseRemove(repo),
	}
}

func (o *MicrocontrollerStoveService) WithTransaction(transaction port_shared.ITransaction) *MicrocontrollerStoveService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *MicrocontrollerStoveService) Get(dtoIn *dto.MicrocontrollerStoveDtoIn) (*dto.MicrocontrollerStoveDtoOut, error) {

	MicrocontrollerStove, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewMicrocontrollerStoveDtoOut()

	dtoOut.Id = MicrocontrollerStove.Id
	dtoOut.IdStove = MicrocontrollerStove.IdStove
	dtoOut.IdMicrocontroller = MicrocontrollerStove.IdMicrocontroller

	if MicrocontrollerStove.CreationDateTime != nil {
		dtoOut.CreationDateTime = MicrocontrollerStove.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if MicrocontrollerStove.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = MicrocontrollerStove.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *MicrocontrollerStoveService) GetAll(conditions ...interface{}) []*dto.MicrocontrollerStoveDtoOut {

	var arrayMicrocontrollerStoveDto []*dto.MicrocontrollerStoveDtoOut

	arrayMicrocontrollerStove := o.ucGetAll.Execute(conditions...)

	for _, MicrocontrollerStove := range arrayMicrocontrollerStove {

		dtoOut := dto.NewMicrocontrollerStoveDtoOut()

		dtoOut.Id = MicrocontrollerStove.Id
		dtoOut.IdStove = MicrocontrollerStove.IdStove
		dtoOut.IdMicrocontroller = MicrocontrollerStove.IdMicrocontroller

		if MicrocontrollerStove.CreationDateTime != nil {
			dtoOut.CreationDateTime = MicrocontrollerStove.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if MicrocontrollerStove.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = MicrocontrollerStove.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayMicrocontrollerStoveDto = append(arrayMicrocontrollerStoveDto, dtoOut)
	}
	return arrayMicrocontrollerStoveDto
}

func (o *MicrocontrollerStoveService) Save(dtoIn *dto.MicrocontrollerStoveDtoIn) error {

	var err error
	MicrocontrollerStove := FactoryMicrocontrollerStove()

	if dtoIn.Id > 0 {
		MicrocontrollerStove.Id = dtoIn.Id
	}

	MicrocontrollerStove.IdStove = dtoIn.IdStove
	MicrocontrollerStove.IdMicrocontroller = dtoIn.IdMicrocontroller

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if MicrocontrollerStove.Id == 0 {
			MicrocontrollerStove.CreationDateTime = &now
		} else {
			MicrocontrollerStoveCurrent, _ := o.ucGet.Execute(MicrocontrollerStove.Id)
			MicrocontrollerStove.CreationDateTime = MicrocontrollerStoveCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		MicrocontrollerStove.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		MicrocontrollerStove.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		MicrocontrollerStove.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(MicrocontrollerStove)
	if err != nil {
		return err
	}
	return nil
}

func (o *MicrocontrollerStoveService) Remove(dtoIn *dto.MicrocontrollerStoveDtoIn) error {

	MicrocontrollerStove := FactoryMicrocontrollerStove()
	if dtoIn.Id > 0 {
		MicrocontrollerStove.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(MicrocontrollerStove)
	if err != nil {
		return err
	}

	return nil
}

func (o *MicrocontrollerStoveService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	// dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
