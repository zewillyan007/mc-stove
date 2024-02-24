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

type StoveService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.StoveIRepository
	ucGet      *usecase.StoveUseCaseGet
	ucSave     *usecase.StoveUseCaseSave
	ucGrid     *usecase.StoveUseCaseGrid
	ucGetAll   *usecase.StoveUseCaseGetAll
	ucRemove   *usecase.StoveUseCaseRemove
}

func NewStoveService(provider port_shared.IResourceProvider) *StoveService {
	repo := types.GetConstructor((*port.StoveIRepository)(nil))(provider.GetDB()).(port.StoveIRepository)
	repo.SetContext(provider.Context())

	return &StoveService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewStoveUseCaseGet(repo),
		ucSave:     usecase.NewStoveUseCaseSave(repo),
		ucGrid:     usecase.NewStoveUseCaseGrid(repo),
		ucGetAll:   usecase.NewStoveUseCaseGetAll(repo),
		ucRemove:   usecase.NewStoveUseCaseRemove(repo),
	}
}

func (o *StoveService) WithTransaction(transaction port_shared.ITransaction) *StoveService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *StoveService) Get(dtoIn *dto.StoveDtoIn) (*dto.StoveDtoOut, error) {

	Stove, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewStoveDtoOut()

	dtoOut.Id = Stove.Id
	dtoOut.Number = Stove.Number
	dtoOut.Length = Stove.Length
	dtoOut.Width = Stove.Width
	dtoOut.Height = Stove.Height

	if Stove.CreationDateTime != nil {
		dtoOut.CreationDateTime = Stove.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if Stove.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Stove.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *StoveService) GetAll(conditions ...interface{}) []*dto.StoveDtoOut {

	var arrayStoveDto []*dto.StoveDtoOut

	arrayStove := o.ucGetAll.Execute(conditions...)

	for _, Stove := range arrayStove {

		dtoOut := dto.NewStoveDtoOut()

		dtoOut.Id = Stove.Id
		dtoOut.Number = Stove.Number
		dtoOut.Length = Stove.Length
		dtoOut.Width = Stove.Width
		dtoOut.Height = Stove.Height

		if Stove.CreationDateTime != nil {
			dtoOut.CreationDateTime = Stove.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if Stove.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Stove.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayStoveDto = append(arrayStoveDto, dtoOut)
	}
	return arrayStoveDto
}

func (o *StoveService) Save(dtoIn *dto.StoveDtoIn) error {

	var err error
	Stove := FactoryStove()

	if dtoIn.Id > 0 {
		Stove.Id = dtoIn.Id
	}

	Stove.Number = dtoIn.Number
	Stove.Length = dtoIn.Length
	Stove.Width = dtoIn.Width
	Stove.Height = dtoIn.Height

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Stove.Id == 0 {
			Stove.CreationDateTime = &now
		} else {
			StoveCurrent, _ := o.ucGet.Execute(Stove.Id)
			Stove.CreationDateTime = StoveCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Stove.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Stove.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Stove.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(Stove)
	if err != nil {
		return err
	}
	return nil
}

func (o *StoveService) Remove(dtoIn *dto.StoveDtoIn) error {

	Stove := FactoryStove()
	if dtoIn.Id > 0 {
		Stove.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(Stove)
	if err != nil {
		return err
	}

	return nil
}

func (o *StoveService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
