package service

import (
	"mc-stove/plantation/core/domain/dto"
	errorsPack "mc-stove/plantation/core/err"
	"mc-stove/plantation/core/port"
	"mc-stove/plantation/core/usecase"
	port_shared "mc-stove/shared/port"
	"mc-stove/shared/types"
	"time"
)

type CurrentParamsService struct {
	// grid.GridCache
	// _cache_         port_shared.ICache
	repository port.CurrentParamsIRepository
	ucGet      *usecase.CurrentParamsUseCaseGet
	ucSave     *usecase.CurrentParamsUseCaseSave
	// ucGrid     *usecase.CurrentParamsUseCaseGrid
	ucGetAll *usecase.CurrentParamsUseCaseGetAll
	ucRemove *usecase.CurrentParamsUseCaseRemove

	//SERVICES
	scMicrocontroller      *MicrocontrollerService
	scMicrocontrollerStove *MicrocontrollerStoveService
	scStove                *StoveService
	scStovePlant           *StovePlantService
	scPlant                *PlantService
}

func NewCurrentParamsService(provider port_shared.IResourceProvider, scMicrocontroller *MicrocontrollerService,
	scMicrocontrollerStove *MicrocontrollerStoveService, scStove *StoveService, scStovePlant *StovePlantService, scPlant *PlantService) *CurrentParamsService {
	repo := types.GetConstructor((*port.CurrentParamsIRepository)(nil))(provider.GetDB()).(port.CurrentParamsIRepository)
	repo.SetContext(provider.Context())

	return &CurrentParamsService{
		// _cache_:         provider.GetCache(),
		repository: repo,
		ucGet:      usecase.NewCurrentParamsUseCaseGet(repo),
		ucSave:     usecase.NewCurrentParamsUseCaseSave(repo),
		// ucGrid:     usecase.NewCurrentParamsUseCaseGrid(repo),
		ucGetAll:               usecase.NewCurrentParamsUseCaseGetAll(repo),
		ucRemove:               usecase.NewCurrentParamsUseCaseRemove(repo),
		scMicrocontroller:      scMicrocontroller,
		scMicrocontrollerStove: scMicrocontrollerStove,
		scStove:                scStove,
		scStovePlant:           scStovePlant,
		scPlant:                scPlant,
	}
}

func (o *CurrentParamsService) WithTransaction(transaction port_shared.ITransaction) *CurrentParamsService {
	o.repository.WithTransaction(transaction)
	return o
}

// func (o *CurrentParamsService) Get(dtoIn *dto.CurrentParamsDtoIn) (*dto.CurrentParamsDtoOut, error) {

// 	CurrentParams, err := o.ucGet.Execute(dtoIn.Id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	dtoOut := dto.NewCurrentParamsDtoOut()

// 	dtoOut.Id = CurrentParams.Id
// 	dtoOut.Type = CurrentParams.Type
// 	dtoOut.SerialNumber = CurrentParams.SerialNumber

// 	if CurrentParams.CreationDateTime != nil {
// 		dtoOut.CreationDateTime = CurrentParams.CreationDateTime.Format("2006-01-02 15:04:05")
// 	}

// 	if CurrentParams.ChangeDateTime != nil {
// 		dtoOut.ChangeDateTime = CurrentParams.ChangeDateTime.Format("2006-01-02 15:04:05")
// 	}

// 	return dtoOut, nil
// }

// func (o *CurrentParamsService) GetAll(conditions ...interface{}) []*dto.CurrentParamsDtoOut {

// 	var arrayCurrentParamsDto []*dto.CurrentParamsDtoOut

// 	arrayCurrentParams := o.ucGetAll.Execute(conditions...)

// 	for _, CurrentParams := range arrayCurrentParams {

// 		dtoOut := dto.NewCurrentParamsDtoOut()

// 		dtoOut.Id = CurrentParams.Id
// 		dtoOut.Type = CurrentParams.Type
// 		dtoOut.SerialNumber = CurrentParams.SerialNumber

// 		if CurrentParams.CreationDateTime != nil {
// 			dtoOut.CreationDateTime = CurrentParams.CreationDateTime.Format("2006-01-02 15:04:05")
// 		}

// 		if CurrentParams.ChangeDateTime != nil {
// 			dtoOut.ChangeDateTime = CurrentParams.ChangeDateTime.Format("2006-01-02 15:04:05")
// 		}

// 		arrayCurrentParamsDto = append(arrayCurrentParamsDto, dtoOut)
// 	}
// 	return arrayCurrentParamsDto
// }

func (o *CurrentParamsService) Save(dtoIn *dto.CurrentParamsDtoIn) error {

	var err error
	CurrentParams := FactoryCurrentParams()

	if dtoIn.Id > 0 {
		CurrentParams.Id = dtoIn.Id
	}

	CurrentParams.Temperature = dtoIn.Temperature
	CurrentParams.Moisture = dtoIn.Moisture

	arrayMicrocontroller := o.scMicrocontroller.GetAll("serial_number = ?", dtoIn.MicroSerialNumber)
	if len(arrayMicrocontroller) == 0 {
		return errorsPack.PlantErrorSpecies
	}

	arrayMicrocontrollerStove := o.scMicrocontrollerStove.GetAll("id_microcontroller = ?", arrayMicrocontroller[0].Id)
	if len(arrayMicrocontrollerStove) == 0 {
		return errorsPack.PlantErrorSpecies
	}

	arrayStove := o.scStove.GetAll("id = ?", arrayMicrocontrollerStove[0].IdStove)
	if len(arrayStove) == 0 {
		return errorsPack.PlantErrorSpecies
	}

	arrayStovePlant := o.scStovePlant.GetAll("id_stove = ?", arrayStove[0].Id)
	if len(arrayStovePlant) == 0 {
		return errorsPack.PlantErrorSpecies
	}

	arrayPlant := o.scPlant.GetAll("id = ?", arrayStovePlant[0].IdPlant)
	if len(arrayPlant) == 0 {
		return errorsPack.PlantErrorSpecies
	}

	CurrentParams.IdStove = arrayStove[0].Id
	CurrentParams.IdPlant = arrayPlant[0].Id
	CurrentParams.StoveNumber = arrayStove[0].Number
	CurrentParams.PlantSpecies = arrayPlant[0].Species

	now := time.Now()

	if len(dtoIn.CreationDateTime) == 0 {
		if CurrentParams.Id == 0 {
			CurrentParams.CreationDateTime = &now
		} else {
			CurrentParamsCurrent, _ := o.ucGet.Execute(CurrentParams.Id)
			CurrentParams.CreationDateTime = CurrentParamsCurrent.CreationDateTime
		}
	} else {
		CreationDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.CreationDateTime)
		if err != nil {
			return err
		}
		CurrentParams.CreationDateTime = &CreationDateTime
	}

	// if len(dtoIn.ChangeDateTime) == 0 {
	// 	CurrentParams.ChangeDateTime = &now
	// } else {
	// 	ChangeDateTime, err := time.Parse("2006-01-02 15:04:05", dtoIn.ChangeDateTime)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	CurrentParams.ChangeDateTime = &ChangeDateTime
	// }

	_, err = o.ucSave.Execute(CurrentParams)
	if err != nil {
		return err
	}
	return nil
}

// func (o *CurrentParamsService) Remove(dtoIn *dto.CurrentParamsDtoIn) error {

// 	CurrentParams := FactoryCurrentParams()
// 	if dtoIn.Id > 0 {
// 		CurrentParams.Id = dtoIn.Id
// 	}
// 	err := o.ucRemove.Execute(CurrentParams)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (o *CurrentParamsService) Grid(GridConfig *grid.GridConfig) (map[string]interface{}, error) {
// 	var dataGrid map[string]interface{}
// 	var err error

// 	// if o._cache_ != nil {
// 	// 	dataGrid, err = o.Cache(o._cache_, GridConfig, o.ucGrid)
// 	// } else {
// 	dataGrid, err = o.ucGrid.Execute(GridConfig)
// 	// }

// 	return dataGrid, err
// }
