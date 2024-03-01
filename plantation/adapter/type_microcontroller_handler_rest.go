package adapter

import (
	"mc-stove/plantation/core/domain/dto"
	"mc-stove/plantation/core/service"
	adapter_shared "mc-stove/shared/adapter"
	"mc-stove/shared/connection/audit"
	"mc-stove/shared/constant"
	"mc-stove/shared/grid"
	"mc-stove/shared/resource"
	"net/http"
	"strconv"
)

type TypeMicrocontrollerHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewTypeMicrocontrollerHandlerRest(resource *resource.ServerResource) *TypeMicrocontrollerHandlerRest {
	return &TypeMicrocontrollerHandlerRest{
		resource: resource,
	}
}

func (h *TypeMicrocontrollerHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_PLANT, h.resource.Herror)

	router := h.resource.DefaultRouter("/type-microncontroller", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *TypeMicrocontrollerHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var plantDtoIn *dto.TypeMicrocontrollerDtoIn
		var plantDtoOut *dto.TypeMicrocontrollerDtoOut

		plantDtoIn = dto.NewTypeMicrocontrollerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &plantDtoIn)
		plantDtoOut, err = service.NewTypeMicrocontrollerService(h.resource.Provider(r)).Get(plantDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_PLANT_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, plantDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *TypeMicrocontrollerHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		plants := service.NewTypeMicrocontrollerService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, plants)
	})
}

func (h *TypeMicrocontrollerHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		plantDtoIn := dto.NewTypeMicrocontrollerDtoIn()
		err = h.resource.Restful.BindDataReq(w, r, &plantDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_PLANT_CREATE, codeErr))
		} else {
			transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

			err = service.NewTypeMicrocontrollerService(h.resource.Provider(r)).WithTransaction(transaction).Save(plantDtoIn)

			if err != nil {
				transaction.Rollback(err)
				codeErr, _ := strconv.Atoi(err.Error())
				err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_PLANT_CREATE, codeErr))
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

func (h *TypeMicrocontrollerHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		TypeMicrocontrollerDtoIn := dto.NewTypeMicrocontrollerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &TypeMicrocontrollerDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

		err = service.NewTypeMicrocontrollerService(h.resource.Provider(r)).WithTransaction(transaction).Save(TypeMicrocontrollerDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_PLANT_SAVE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *TypeMicrocontrollerHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		TypeMicrocontrollerDtoIn := dto.NewTypeMicrocontrollerDtoIn()
		h.resource.Restful.BindDataReq(w, r, &TypeMicrocontrollerDtoIn)
		err = service.NewTypeMicrocontrollerService(h.resource.Provider(r)).Remove(TypeMicrocontrollerDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_PLANT_REMOVE, codeErr))
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *TypeMicrocontrollerHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GridConfig := grid.NewGridConfig()
		h.resource.Restful.BindDataReq(w, r, &GridConfig)
		GridConfig = h.GridConfigData(GridConfig)
		dataGrid, err := service.NewTypeMicrocontrollerService(h.resource.Provider(r)).Grid(GridConfig)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_PLANT_GRID, codeErr))
		} else {
			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
			} else {
				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
			}
		}
	})
}
