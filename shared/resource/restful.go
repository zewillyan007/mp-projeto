package resource

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	port_shared "mp-projeto/shared/port"
	"mp-projeto/shared/presentation"

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

func (o *Restful) Logger() {

	if o.log != nil {
		params := string(o.content)
		params = strings.ReplaceAll(params, "\n", "")
		params = strings.ReplaceAll(params, "\t", "")
		o.log.SetExtraPart("reqid", o.request.Header.Get("reqid")).SetExtraPart("method", o.request.Method).SetExtraPart("url", o.request.URL.String()).Info(params)
	}
}

func (o *Restful) GetContent() []byte {
	return o.content
}

func (o *Restful) AddHeader(key string, value string) {
	o.response.Header().Add(key, value)
}

func (o *Restful) Write(data []byte) {
	o.response.Write(data)
}

func (o *Restful) GetRequestHeaderValue(key string) string {

	var value string
	for _, val := range o.request.Header[key] {
		value = val
	}
	return value
}

func (o *Restful) LoadData(response http.ResponseWriter, request *http.Request) *Restful {

	o.request = request
	o.response = response
	defer o.request.Body.Close()

	switch o.request.Method {
	case "GET", "DELETE":
		field := mux.Vars(o.request)
		reference := make(map[string]interface{})
		for key, value := range field {
			reference[key] = value
		}
		query := o.request.URL.Query()
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
		o.content, _ = json.Marshal(reference)
	case "PUT", "POST":
		switch o.GetRequestHeaderValue("Content-Type") {
		case APPLICATION_FORM:
			o.request.ParseForm()
			content := make(map[string]string)
			for key, values := range o.request.Form {
				for _, value := range values {
					content[key] = value
				}
			}
			o.content, _ = json.Marshal(content)
		case APPLICATION_JSON:
			o.content, _ = ioutil.ReadAll(o.request.Body)
			o.request.Body = ioutil.NopCloser(bytes.NewBuffer(o.content))
			if o.request.Method == "PUT" {
				field := mux.Vars(o.request)
				reference := make(map[string]interface{})
				json.Unmarshal(o.content, &reference)
				reference["id"] = field["id"]
				o.content, _ = json.Marshal(reference)
			}
		}
	}
	o.Logger()
	return o
}

func (o *Restful) BindData(reference interface{}) error {

	return json.Unmarshal(o.content, reference)
}

func (o *Restful) GenericStruct() map[string]interface{} {

	reference := make(map[string]interface{})
	json.Unmarshal(o.content, &reference)
	return reference
}
