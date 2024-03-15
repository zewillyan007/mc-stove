package adapter

import (
	"mc-stove/plantation/core/domain/dto"
	"mc-stove/plantation/core/service"
	adapter_shared "mc-stove/shared/adapter"
	"mc-stove/shared/connection/audit"
	"mc-stove/shared/constant"
	"mc-stove/shared/resource"
	"net/http"
	"strconv"
)

type CurrentParamsHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewCurrentParamsHandlerRest(resource *resource.ServerResource) *CurrentParamsHandlerRest {
	return &CurrentParamsHandlerRest{
		resource: resource,
	}
}

func (h *CurrentParamsHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_PLANT, h.resource.Herror)

	router := h.resource.DefaultRouter("/current-params", true)
	// router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	// router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	// router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	// router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	// router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

// func (h *CurrentParamsHandlerRest) get() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var err error
// 		var currentParamsDtoIn *dto.CurrentParamsDtoIn
// 		var currentParamsDtoOut *dto.CurrentParamsDtoOut

// 		currentParamsDtoIn = dto.NewCurrentParamsDtoIn()
// 		h.resource.Restful.BindDataReq(w, r, &currentParamsDtoIn)
// 		currentParamsDtoOut, err = service.NewCurrentParamsService(h.resource.Provider(r)).Get(currentParamsDtoIn)

// 		if err != nil {
// 			codeErr, _ := strconv.Atoi(err.Error())
// 			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_CURRENT_PARAMS_GET, codeErr))
// 		} else {
// 			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, currentParamsDtoOut)
// 		}

// 		if err != nil {
// 			h.resource.Log.Error("%s\n", err.Error())
// 		}
// 	})
// }

// func (h *CurrentParamsHandlerRest) getAll() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		currentParams := service.NewCurrentParamsService(h.resource.Provider(r)).GetAll()
// 		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, currentParams)
// 	})
// }

func (h *CurrentParamsHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		scMicrocontroller := service.NewMicrocontrollerService(h.resource.Provider(r))
		scMicrocontrollerStove := service.NewMicrocontrollerStoveService(h.resource.Provider(r))
		scStove := service.NewStoveService(h.resource.Provider(r))
		scStovePlant := service.NewStovePlantService(h.resource.Provider(r))
		scPlant := service.NewPlantService(h.resource.Provider(r))

		currentParamsDtoIn := dto.NewCurrentParamsDtoIn()
		err = h.resource.Restful.BindDataReq(w, r, &currentParamsDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_CURRENT_PARAMS_CREATE, codeErr))
		} else {
			transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

			err = service.NewCurrentParamsService(h.resource.Provider(r), scMicrocontroller, scMicrocontrollerStove, scStove, scStovePlant,
				scPlant).WithTransaction(transaction).Save(currentParamsDtoIn)

			if err != nil {
				transaction.Rollback(err)
				codeErr, _ := strconv.Atoi(err.Error())
				err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_CURRENT_PARAMS_CREATE, codeErr))
			} else {
				transaction.Commit()
				err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
			}
		}
		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

// func (h *CurrentParamsHandlerRest) save() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var err error

// 		CurrentParamsDtoIn := dto.NewCurrentParamsDtoIn()
// 		h.resource.Restful.BindDataReq(w, r, &CurrentParamsDtoIn)
// 		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

// 		err = service.NewCurrentParamsService(h.resource.Provider(r)).WithTransaction(transaction).Save(CurrentParamsDtoIn)

// 		if err != nil {
// 			transaction.Rollback(err)
// 			codeErr, _ := strconv.Atoi(err.Error())
// 			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_PLANT_SAVE, codeErr))
// 		} else {
// 			transaction.Commit()
// 			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
// 		}

// 		if err != nil {
// 			h.resource.Log.Error("%s\n", err.Error())
// 		}
// 	})
// }

// func (h *CurrentParamsHandlerRest) remove() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		var err error

// 		CurrentParamsDtoIn := dto.NewCurrentParamsDtoIn()
// 		h.resource.Restful.BindDataReq(w, r, &CurrentParamsDtoIn)
// 		err = service.NewCurrentParamsService(h.resource.Provider(r)).Remove(CurrentParamsDtoIn)

// 		if err != nil {
// 			codeErr, _ := strconv.Atoi(err.Error())
// 			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_PLANT_REMOVE, codeErr))
// 		} else {
// 			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
// 		}

// 		if err != nil {
// 			h.resource.Log.Error("%s\n", err.Error())
// 		}
// 	})
// }

// func (h *CurrentParamsHandlerRest) grid() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		GridConfig := grid.NewGridConfig()
// 		h.resource.Restful.BindDataReq(w, r, &GridConfig)
// 		GridConfig = h.GridConfigData(GridConfig)
// 		dataGrid, err := service.NewCurrentParamsService(h.resource.Provider(r)).Grid(GridConfig)

// 		if err != nil {
// 			codeErr, _ := strconv.Atoi(err.Error())
// 			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_PLANT_GRID, codeErr))
// 		} else {
// 			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
// 				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
// 			} else {
// 				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
// 			}
// 		}
// 	})
// }
