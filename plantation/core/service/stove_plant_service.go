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

type StovePlantService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.StovePlantIRepository
	ucGet      *usecase.StovePlantUseCaseGet
	ucSave     *usecase.StovePlantUseCaseSave
	// ucGrid     *usecase.StovePlantUseCaseGrid
	ucGetAll *usecase.StovePlantUseCaseGetAll
	ucRemove *usecase.StovePlantUseCaseRemove
}

func NewStovePlantService(provider port_shared.IResourceProvider) *StovePlantService {
	repo := types.GetConstructor((*port.StovePlantIRepository)(nil))(provider.GetDB()).(port.StovePlantIRepository)
	repo.SetContext(provider.Context())

	return &StovePlantService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewStovePlantUseCaseGet(repo),
		ucSave:     usecase.NewStovePlantUseCaseSave(repo),
		// ucGrid:     usecase.NewStovePlantUseCaseGrid(repo),
		ucGetAll: usecase.NewStovePlantUseCaseGetAll(repo),
		ucRemove: usecase.NewStovePlantUseCaseRemove(repo),
	}
}

func (o *StovePlantService) WithTransaction(transaction port_shared.ITransaction) *StovePlantService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *StovePlantService) Get(dtoIn *dto.StovePlantDtoIn) (*dto.StovePlantDtoOut, error) {

	StovePlant, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewStovePlantDtoOut()

	dtoOut.Id = StovePlant.Id
	dtoOut.IdStove = StovePlant.IdStove
	dtoOut.IdPlant = StovePlant.IdPlant

	if StovePlant.CreationDateTime != nil {
		dtoOut.CreationDateTime = StovePlant.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if StovePlant.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = StovePlant.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *StovePlantService) GetAll(conditions ...interface{}) []*dto.StovePlantDtoOut {

	var arrayStovePlantDto []*dto.StovePlantDtoOut

	arrayStovePlant := o.ucGetAll.Execute(conditions...)

	for _, StovePlant := range arrayStovePlant {

		dtoOut := dto.NewStovePlantDtoOut()

		dtoOut.Id = StovePlant.Id
		dtoOut.IdStove = StovePlant.IdStove
		dtoOut.IdPlant = StovePlant.IdPlant

		if StovePlant.CreationDateTime != nil {
			dtoOut.CreationDateTime = StovePlant.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if StovePlant.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = StovePlant.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayStovePlantDto = append(arrayStovePlantDto, dtoOut)
	}
	return arrayStovePlantDto
}

func (o *StovePlantService) Save(dtoIn *dto.StovePlantDtoIn) error {

	var err error
	StovePlant := FactoryStovePlant()

	if dtoIn.Id > 0 {
		StovePlant.Id = dtoIn.Id
	}

	StovePlant.IdStove = dtoIn.IdStove
	StovePlant.IdPlant = dtoIn.IdPlant

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if StovePlant.Id == 0 {
			StovePlant.CreationDateTime = &now
		} else {
			StovePlantCurrent, _ := o.ucGet.Execute(StovePlant.Id)
			StovePlant.CreationDateTime = StovePlantCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		StovePlant.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		StovePlant.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		StovePlant.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(StovePlant)
	if err != nil {
		return err
	}
	return nil
}

func (o *StovePlantService) Remove(dtoIn *dto.StovePlantDtoIn) error {

	StovePlant := FactoryStovePlant()
	if dtoIn.Id > 0 {
		StovePlant.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(StovePlant)
	if err != nil {
		return err
	}

	return nil
}

func (o *StovePlantService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	// dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
