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

type StoveHandlerRest struct {
	adapter_shared.Handler
	resource *resource.ServerResource
}

func NewStoveHandlerRest(resource *resource.ServerResource) *StoveHandlerRest {
	return &StoveHandlerRest{
		resource: resource,
	}
}

func (h *StoveHandlerRest) MakeRoutes() {

	h.ConfigError(constant.HDR_STOVE, h.resource.Herror)

	router := h.resource.DefaultRouter("/stoves", true)
	router.Handle("", h.getAll()).Methods(http.MethodGet)
	router.Handle("", h.create()).Methods(http.MethodPost)
	router.Handle("/{id:[0-9]+}", h.get()).Methods(http.MethodGet)
	router.Handle("/{id:[0-9]+}", h.save()).Methods(http.MethodPut)
	router.Handle("/{id:[0-9]+}", h.remove()).Methods(http.MethodDelete)
	router.Handle("/grid", h.grid()).Methods(http.MethodPost)
}

func (h *StoveHandlerRest) get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var stoveDtoIn *dto.StoveDtoIn
		var stoveDtoOut *dto.StoveDtoOut

		stoveDtoIn = dto.NewStoveDtoIn()
		h.resource.Restful.BindDataReq(w, r, &stoveDtoIn)
		stoveDtoOut, err = service.NewStoveService(h.resource.Provider(r)).Get(stoveDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_STOVE_GET, codeErr))
		} else {
			err = h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, stoveDtoOut)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *StoveHandlerRest) getAll() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		stoves := service.NewStoveService(h.resource.Provider(r)).GetAll()
		h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, stoves)
	})
}

func (h *StoveHandlerRest) create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		stoveDtoIn := dto.NewStoveDtoIn()
		h.resource.Restful.BindDataReq(w, r, &stoveDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Insert)

		err = service.NewStoveService(h.resource.Provider(r)).WithTransaction(transaction).Save(stoveDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_STOVE_CREATE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *StoveHandlerRest) save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		StoveDtoIn := dto.NewStoveDtoIn()
		h.resource.Restful.BindDataReq(w, r, &StoveDtoIn)
		transaction := adapter_shared.BeginTrans(h.resource.Provider(r), audit.Update)

		err = service.NewStoveService(h.resource.Provider(r)).WithTransaction(transaction).Save(StoveDtoIn)

		if err != nil {
			transaction.Rollback(err)
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_STOVE_SAVE, codeErr))
		} else {
			transaction.Commit()
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *StoveHandlerRest) remove() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error

		StoveDtoIn := dto.NewStoveDtoIn()
		h.resource.Restful.BindDataReq(w, r, &StoveDtoIn)
		err = service.NewStoveService(h.resource.Provider(r)).Remove(StoveDtoIn)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			err = h.resource.Restful.ResponseError(w, r, resource.HTTP_NOT_FOUND, h.Lerr(constant.HDR_EPT_STOVE_REMOVE, codeErr))
		} else {
			err = h.resource.Restful.Response(w, r, resource.HTTP_OK)
		}

		if err != nil {
			h.resource.Log.Error("%s\n", err.Error())
		}
	})
}

func (h *StoveHandlerRest) grid() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		GridConfig := grid.NewGridConfig()
		h.resource.Restful.BindDataReq(w, r, &GridConfig)
		GridConfig = h.GridConfigData(GridConfig)
		dataGrid, err := service.NewStoveService(h.resource.Provider(r)).Grid(GridConfig)

		if err != nil {
			codeErr, _ := strconv.Atoi(err.Error())
			h.resource.Restful.ResponseError(w, r, resource.HTTP_BAD_REQUEST, h.Lerr(constant.HDR_EPT_STOVE_GRID, codeErr))
		} else {
			if GridConfig.Export != nil && len(GridConfig.Export.Value) > 0 {
				grid.ResponseDataGrid(w, GridConfig.Export.Type, dataGrid, "product")
			} else {
				h.resource.Restful.ResponseData(w, r, resource.HTTP_OK, dataGrid)
			}
		}
	})
}
