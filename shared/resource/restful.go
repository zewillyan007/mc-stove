package resource

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	port_shared "mc-stove/shared/port"
	"mc-stove/shared/presentation"

	"github.com/gorilla/mux"
)

const HTTP_OK int = 200
const HTTP_CREATED int = 201
const HTTP_NO_CONTENT int = 204
const HTTP_FOUND int = 302
const HTTP_NOT_MODIFIED int = 304
const HTTP_BAD_REQUEST int = 400
const HTTP_UNAUTHORIZED int = 401
const HTTP_FORBIDDEN int = 403
const HTTP_NOT_FOUND int = 404
const HTTP_INTERNAL_SERVER_ERROR int = 500
const APPLICATION_FORM string = "application/x-www-form-urlencoded"
const APPLICATION_JSON string = "application/json"

type Restful struct {
	content  []byte
	log      port_shared.ILogger
	request  *http.Request
	response http.ResponseWriter
	*presentation.ResponsePattern
}

func NewRestful(Logger ...port_shared.ILogger) *Restful {
	Restful := &Restful{
		content:  []byte{},
		log:      nil,
		request:  &http.Request{},
		response: nil,
	}

	if len(Logger) > 0 {
		Restful.log = Logger[0]
		if len(Logger) > 1 {
			Restful.ResponsePattern = presentation.NewResponsePattern(Logger[1])
		} else {
			Restful.ResponsePattern = presentation.NewResponsePattern()
		}
	} else {
		Restful.ResponsePattern = presentation.NewResponsePattern()
	}

	return Restful
}

func (o *Restful) Logger(content []byte, r *http.Request) {

	if o.log != nil {
		params := string(content)
		params = strings.ReplaceAll(params, "\n", "")
		params = strings.ReplaceAll(params, "\t", "")
		o.log.SetExtraPart("reqid", r.Header.Get("reqid")).SetExtraPart("method", r.Method).SetExtraPart("url", r.URL.String()).Info(params)
	}
}

/*
func (o *Restful) GetContent() []byte {
	return o.content
}

func (o *Restful) AddHeader(key string, value string) {
	o.response.Header().Add(key, value)
}

func (o *Restful) Write(data []byte) {
	o.response.Write(data)
}
*/

func (o *Restful) GetRequestHeaderValue(key string, r *http.Request) string {

	var value string
	for _, val := range r.Header[key] {
		value = val
	}
	return value
}

func (o *Restful) LoadData(response http.ResponseWriter, request *http.Request) []byte {

	var content []byte
	defer request.Body.Close()

	switch request.Method {
	case "GET", "DELETE":
		field := mux.Vars(request)
		reference := make(map[string]interface{})
		for key, value := range field {
			reference[key] = value
		}
		query := request.URL.Query()
		params, _ := url.ParseQuery(query.Encode())
		slice := []string{}
		for key, values := range params {
			_, exists := reference[key]
			if !exists {
				if len(values) > 1 {
					for _, value := range values {
						slice = append(slice, value)
					}
					reference[key] = strings.Join(slice, ",")
				} else {
					for key, value := range params {
						reference[key] = strings.Join(value, "")
					}
				}
			}
		}
		content, _ = json.Marshal(reference)
	case "PUT", "POST":
		switch o.GetRequestHeaderValue("Content-Type", request) {
		case APPLICATION_FORM:
			request.ParseForm()
			contentMap := make(map[string]string)
			for key, values := range request.Form {
				for _, value := range values {
					contentMap[key] = value
				}
			}
			content, _ = json.Marshal(contentMap)
		case APPLICATION_JSON:
			content, _ = ioutil.ReadAll(request.Body)
			request.Body = ioutil.NopCloser(bytes.NewBuffer(content))
			if request.Method == "PUT" {
				field := mux.Vars(request)
				reference := make(map[string]interface{})
				json.Unmarshal(content, &reference)
				reference["id"] = field["id"]
				content, _ = json.Marshal(reference)
			}
		}
	}
	o.Logger(content, request)
	return content
}

func (o *Restful) BindData(reference interface{}, content []byte) error {

	return json.Unmarshal(content, reference)
}

func (o *Restful) BindDataReq(response http.ResponseWriter, request *http.Request, reference interface{}) error {

	content := o.LoadData(response, request)
	return o.BindData(reference, content)
}

func (o *Restful) GenericStruct(content []byte) map[string]interface{} {

	reference := make(map[string]interface{})
	json.Unmarshal(content, &reference)
	return reference
}
