package presentation

import (
	"encoding/json"
	port_shared "mc-stove/shared/port"
	"net/http"
)

type RepStatus struct {
	Status int `json:"status"`
}

type RepData struct {
	Data interface{} `json:"data"`
}

type RepErr struct {
	Error interface{} `json:"err"`
}

type RepMessage struct {
	Message interface{} `json:"msg"`
}

type RepResponse struct {
	*RepStatus
}

func NewRepResponse() *RepResponse {
	return &RepResponse{
		RepStatus: &RepStatus{
			Status: 0,
		},
	}
}

type RepResponseData struct {
	*RepStatus
	*RepData
}

func NewRepResponseData() *RepResponseData {
	return &RepResponseData{
		RepStatus: &RepStatus{
			Status: 0,
		},
		RepData: &RepData{
			Data: nil,
		},
	}
}

type RepResponseError struct {
	*RepStatus
	*RepErr
}

func NewRepResponseError() *RepResponseError {
	return &RepResponseError{
		RepStatus: &RepStatus{
			Status: 0,
		},
		RepErr: &RepErr{
			Error: nil,
		},
	}
}

type RepResponseErrorData struct {
	*RepStatus
	*RepErr
	*RepData
}

func NewRepResponseErrorData() *RepResponseErrorData {
	return &RepResponseErrorData{
		RepStatus: &RepStatus{
			Status: 0,
		},
		RepErr: &RepErr{
			Error: nil,
		},
		RepData: &RepData{
			Data: nil,
		},
	}
}

type RepResponseMessage struct {
	*RepStatus
	*RepMessage
}

func NewRepResponseMessage() *RepResponseMessage {
	return &RepResponseMessage{
		RepStatus: &RepStatus{Status: 0},
		RepMessage: &RepMessage{
			Message: nil,
		},
	}
}

type RepResponseMessageData struct {
	*RepStatus
	*RepMessage
	*RepData
}

func NewRepResponseMessageData() *RepResponseMessageData {
	return &RepResponseMessageData{
		RepStatus: &RepStatus{
			Status: 0,
		},
		RepMessage: &RepMessage{
			Message: nil,
		},
		RepData: &RepData{
			Data: nil,
		},
	}
}

type ResponsePattern struct {
	log port_shared.ILogger
}

func (o *ResponsePattern) _write(w http.ResponseWriter, r *http.Request, bytes []byte) {

	if o.log != nil {
		o.log.SetExtraPart("reqid", r.Header.Get("reqid")).SetExtraPart("method", r.Method).SetExtraPart("url", r.URL.String()).Info(string(bytes))
	}
	w.Write(bytes)
}

func (o *ResponsePattern) send(w http.ResponseWriter, r *http.Request, msg interface{}) error {

	ret, _err := json.Marshal(msg)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func NewResponsePattern(Logger ...port_shared.ILogger) *ResponsePattern {

	ResponsePattern := &ResponsePattern{}
	if len(Logger) > 0 {
		ResponsePattern.log = Logger[0]
	}
	return ResponsePattern
}

func (o *ResponsePattern) Response(w http.ResponseWriter, r *http.Request, status int) error {

	w.WriteHeader(status)
	msg := NewRepResponse()
	msg.Status = status
	return o.send(w, r, msg)
}

func (o *ResponsePattern) ResponseData(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {

	w.WriteHeader(status)
	msg := NewRepResponseData()
	msg.Status = status
	msg.Data = data
	return o.send(w, r, msg)
}

func (o *ResponsePattern) ResponseError(w http.ResponseWriter, r *http.Request, status int, err interface{}) error {

	w.WriteHeader(status)
	msg := NewRepResponseError()
	msg.Status = status
	msg.Error = err
	return o.send(w, r, msg)
}

func (o *ResponsePattern) ResponseErrorData(w http.ResponseWriter, r *http.Request, status int, err, data interface{}) error {

	w.WriteHeader(status)
	msg := NewRepResponseErrorData()
	msg.Status = status
	msg.Error = err
	msg.Data = data
	return o.send(w, r, msg)
}

func (o *ResponsePattern) ResponseMessage(w http.ResponseWriter, r *http.Request, status int, message interface{}) error {

	w.WriteHeader(status)
	msg := NewRepResponseMessage()
	msg.Status = status
	msg.Message = message
	return o.send(w, r, msg)
}

func (o *ResponsePattern) ResponseMessageData(w http.ResponseWriter, r *http.Request, status int, message, data interface{}) error {

	w.WriteHeader(status)
	msg := NewRepResponseMessageData()
	msg.Status = status
	msg.Message = message
	msg.Data = message
	return o.send(w, r, msg)
}
