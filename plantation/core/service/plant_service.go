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

type PlantService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.PlantIRepository
	ucGet      *usecase.PlantUseCaseGet
	ucSave     *usecase.PlantUseCaseSave
	ucGrid     *usecase.PlantUseCaseGrid
	ucGetAll   *usecase.PlantUseCaseGetAll
	ucRemove   *usecase.PlantUseCaseRemove
}

func NewPlantService(provider port_shared.IResourceProvider) *PlantService {
	repo := types.GetConstructor((*port.PlantIRepository)(nil))(provider.GetDB()).(port.PlantIRepository)
	repo.SetContext(provider.Context())

	return &PlantService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewPlantUseCaseGet(repo),
		ucSave:     usecase.NewPlantUseCaseSave(repo),
		ucGrid:     usecase.NewPlantUseCaseGrid(repo),
		ucGetAll:   usecase.NewPlantUseCaseGetAll(repo),
		ucRemove:   usecase.NewPlantUseCaseRemove(repo),
	}
}

func (o *PlantService) WithTransaction(transaction port_shared.ITransaction) *PlantService {
	o.repository.WithTransaction(transaction)
	return o
}

func (o *PlantService) Get(dtoIn *dto.PlantDtoIn) (*dto.PlantDtoOut, error) {

	Plant, err := o.ucGet.Execute(dtoIn.Id)
	if err != nil {
		return nil, err
	}

	dtoOut := dto.NewPlantDtoOut()

	dtoOut.Id = Plant.Id
	dtoOut.Species = Plant.Species
	dtoOut.Temperature = Plant.Temperature
	dtoOut.Ph = Plant.Ph
	dtoOut.Moisture = Plant.Moisture

	if Plant.CreationDateTime != nil {
		dtoOut.CreationDateTime = Plant.CreationDateTime.Format("2006-01-02 15:04:05")
	}

	if Plant.ChangeDateTime != nil {
		dtoOut.ChangeDateTime = Plant.ChangeDateTime.Format("2006-01-02 15:04:05")
	}

	return dtoOut, nil
}

func (o *PlantService) GetAll(conditions ...interface{}) []*dto.PlantDtoOut {

	var arrayPlantDto []*dto.PlantDtoOut

	arrayPlant := o.ucGetAll.Execute(conditions...)

	for _, Plant := range arrayPlant {

		dtoOut := dto.NewPlantDtoOut()

		dtoOut.Id = Plant.Id
		dtoOut.Species = Plant.Species
		dtoOut.Temperature = Plant.Temperature
		dtoOut.Ph = Plant.Ph
		dtoOut.Moisture = Plant.Moisture

		if Plant.CreationDateTime != nil {
			dtoOut.CreationDateTime = Plant.CreationDateTime.Format("2006-01-02 15:04:05")
		}

		if Plant.ChangeDateTime != nil {
			dtoOut.ChangeDateTime = Plant.ChangeDateTime.Format("2006-01-02 15:04:05")
		}

		arrayPlantDto = append(arrayPlantDto, dtoOut)
	}
	return arrayPlantDto
}

func (o *PlantService) Save(dtoIn *dto.PlantDtoIn) error {

	var err error
	Plant := FactoryPlant()

	if dtoIn.Id > 0 {
		Plant.Id = dtoIn.Id
	}

	Plant.Species = dtoIn.Species
	Plant.Temperature = dtoIn.Temperature
	Plant.Ph = dtoIn.Ph
	Plant.Moisture = dtoIn.Moisture

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if Plant.Id == 0 {
			Plant.CreationDateTime = &now
		} else {
			PlantCurrent, _ := o.ucGet.Execute(Plant.Id)
			Plant.CreationDateTime = PlantCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		Plant.CreationDateTime = &CreationDateTime
	}

	if len(dtoIn.ChangeDateTime) == 0 {
		Plant.ChangeDateTime = &now
	} else {
		ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
		if err != nil {
			return err
		}
		Plant.ChangeDateTime = &ChangeDateTime
	}

	_, err = o.ucSave.Execute(Plant)
	if err != nil {
		return err
	}
	return nil
}

func (o *PlantService) Remove(dtoIn *dto.PlantDtoIn) error {

	Plant := FactoryPlant()
	if dtoIn.Id > 0 {
		Plant.Id = dtoIn.Id
	}
	err := o.ucRemove.Execute(Plant)
	if err != nil {
		return err
	}

	return nil
}

func (o *PlantService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
	var dataGrid map[string]interface{}
	var err error

	// if o._cache_ != nil {
	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
	// } else {
	dataGrid, err = o.ucGrid.Execute(GridConfig)
	// }

	return dataGrid, err
}
