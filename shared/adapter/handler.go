package adapter

import (
	"fmt"
	"mc-stove/shared/grid"
	"mc-stove/shared/herror"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	IdHandler string
	herror    *herror.HandlerError
}

type HandlerError struct {
	Err struct {
		Cod string `json:"cod"`
		Cer string `json:"cer"`
		Msg string `json:"msg"`
	} `json:"err"`
}

type HandlerErrorData struct {
	Err struct {
		Cod  string      `json:"cod"`
		Cer  string      `json:"cer"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	} `json:"err"`
}

func (o *Handler) ConfigError(IdHandler string, herror *herror.HandlerError) {

	o.IdHandler = IdHandler
	o.herror = herror
}

func (o *Handler) prepareError(idEndPoint string) HandlerError {

	errh := HandlerError{}
	errh.Err.Cod = o.IdHandler + "-" + idEndPoint
	return errh
}

func (o *Handler) prepareErrorData(idEndPoint string) HandlerErrorData {

	errh := HandlerErrorData{}
	errh.Err.Cod = o.IdHandler + "-" + idEndPoint
	return errh
}

func (o *Handler) genericError(idEndPoint string, id int) map[string]interface{} {

	var ccode, ecode string
	_error_ := make(map[string]interface{})
	clenth := o.herror.Dictionary.DigitsClass
	dlenth := o.herror.Dictionary.DigitsError

	if id <= 0 {

		formatc := fmt.Sprintf("%v", clenth)
		formate := fmt.Sprintf("%v", dlenth)
		ccode = fmt.Sprintf("%0"+formatc+"d", 0)
		ecode = fmt.Sprintf("%0"+formate+"d", 0)

	} else {

		digits := o.herror.SplitInt(id)
		formatc := fmt.Sprintf("%v", clenth)
		formate := fmt.Sprintf("%v", dlenth)
		ccode = fmt.Sprintf("%0"+formatc+"d", o.herror.SliceToInt(digits[:len(digits)-dlenth]))
		ecode = fmt.Sprintf("%0"+formate+"d", o.herror.SliceToInt(digits[len(digits)-dlenth:]))
	}

	_error_["typ"] = "U"
	_error_["cls"] = "Unmapped"
	_error_["cod"] = o.IdHandler + "-" + idEndPoint + "-E" + ccode + ecode
	_error_["cer"] = ccode + ecode
	_error_["msg"] = "Error Unmapped"
	return _error_
}

func (o *Handler) Gerr(idEndPoint string, id int) map[string]interface{} {

	_error_ := make(map[string]interface{})
	errh := o.prepareError(idEndPoint)
	erri := o.herror.Error(id)
	if erri != nil {
		_error_["typ"] = "G"
		_error_["cls"] = erri.Clt
		_error_["cod"] = errh.Err.Cod + "-" + erri.Cod
		_error_["cer"] = strings.ReplaceAll(erri.Cod, "E", "")
		_error_["msg"] = erri.Clt + ": " + erri.Msg
	} else {
		_error_ = o.genericError(idEndPoint, id)
	}
	return _error_
}

func (o *Handler) GerrData(idEndPoint string, id int, data ...interface{}) map[string]interface{} {

	_error_ := make(map[string]interface{})
	errh := o.prepareErrorData(idEndPoint)
	erri := o.herror.Error(id)
	if erri != nil {
		_error_["typ"] = "G"
		_error_["cls"] = erri.Clt
		_error_["cod"] = errh.Err.Cod + "-" + erri.Cod
		_error_["cer"] = strings.ReplaceAll(erri.Cod, "E", "")
		_error_["msg"] = erri.Clt + ": " + erri.Msg
		if len(data) > 0 {
			_error_["data"] = data[0]
		}
	} else {
		_error_ = o.genericError(idEndPoint, id)
	}
	return _error_
}

func (o *Handler) Lerr(idEndPoint string, id int) map[string]interface{} {

	_error_ := make(map[string]interface{})
	errh := o.prepareError(idEndPoint)
	erri := o.herror.ErrorApp(id)
	if erri != nil {
		_error_["typ"] = "A"
		_error_["cls"] = erri.Clt
		_error_["cod"] = errh.Err.Cod + "-" + erri.Cod
		_error_["cer"] = strings.ReplaceAll(erri.Cod, "E", "")
		_error_["msg"] = erri.Msg
	} else {
		_error_ = o.genericError(idEndPoint, id)
	}
	return _error_
}

func (o *Handler) LerrData(idEndPoint string, id int, data ...interface{}) map[string]interface{} {

	_error_ := make(map[string]interface{})
	errh := o.prepareErrorData(idEndPoint)
	erri := o.herror.ErrorApp(id)
	if erri != nil {
		_error_["typ"] = "A"
		_error_["cls"] = erri.Clt
		_error_["cod"] = errh.Err.Cod + "-" + erri.Cod
		_error_["cer"] = strings.ReplaceAll(erri.Cod, "E", "")
		_error_["msg"] = erri.Msg
		if len(data) > 0 {
			_error_["data"] = data[0]
		}
	} else {
		_error_ = o.genericError(idEndPoint, id)
	}

	return _error_
}

func (o *Handler) GridConfigData(GridConfig *grid.GridConfig) *grid.GridConfig {

	if GridConfig.Export != nil {
		if GridConfig.Export.Value == "all_pages" {
			GridConfig.Page = "1"
			GridConfig.RowsPage = "1000000000"
		}
		if (GridConfig.Export != nil) && len(GridConfig.Export.Type) == 0 {
			GridConfig.Export.Type = "csv"
		}
	}

	return GridConfig
}

func (o *Handler) CurrentRouteStr(r *http.Request, ix ...uint8) string {
	s, _ := mux.CurrentRoute(r).GetPathTemplate()
	if len(ix) == 0 {
		return s
	} else {
		return strings.Split(s, "/")[ix[0]]
	}
}

func (o *Handler) RootRouteStr(r *http.Request) string {
	return o.CurrentRouteStr(r, 1)
}
