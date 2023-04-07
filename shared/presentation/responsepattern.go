package presentation

import (
	"encoding/json"
	port_shared "mp-projeto/shared/port"
	"net/http"
)

type _status struct {
	Status int `json:"status"`
}

type _data struct {
	Data interface{} `json:"data"`
}

type _err struct {
	Error interface{} `json:"err"`
}

type _message struct {
	Message interface{} `json:"msg"`
}

type _response struct {
	*_status
}

type _responseData struct {
	*_status
	*_data
}

type _responseError struct {
	*_status
	*_err
}

type _responseErrorData struct {
	*_status
	*_err
	*_data
}

type _responseMessage struct {
	*_status
	*_message
}

type _responseMessageData struct {
	*_status
	*_message
	*_data
}

type ResponsePattern struct {
	log port_shared.ILogger
	*_response
	*_responseData
	*_responseError
	*_responseErrorData
	*_responseMessage
	*_responseMessageData
}

func (o *ResponsePattern) _write(w http.ResponseWriter, r *http.Request, bytes []byte) {

	if o.log != nil {
		o.log.SetExtraPart("reqid", r.Header.Get("reqid")).SetExtraPart("method", r.Method).SetExtraPart("url", r.URL.String()).Info(string(bytes))
	}
	w.Write(bytes)
}

func NewResponsePattern(Logger ...port_shared.ILogger) *ResponsePattern {

	ResPattern := &ResponsePattern{
		_response: &_response{
			_status: &_status{},
		},
		_responseData: &_responseData{
			_status: &_status{},
			_data:   &_data{},
		},
		_responseError: &_responseError{
			_status: &_status{},
			_err:    &_err{},
		},
		_responseErrorData: &_responseErrorData{
			_status: &_status{},
			_err:    &_err{},
			_data:   &_data{},
		},
		_responseMessage: &_responseMessage{
			_status:  &_status{},
			_message: &_message{},
		},
		_responseMessageData: &_responseMessageData{
			_status:  &_status{},
			_message: &_message{},
			_data:    &_data{},
		},
	}
	if len(Logger) > 0 {
		ResPattern.log = Logger[0]
	}
	return ResPattern
}

func (o *ResponsePattern) Response(w http.ResponseWriter, r *http.Request, status int) error {

	w.WriteHeader(status)
	o._response.Status = status
	ret, _err := json.Marshal(o._response)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func (o *ResponsePattern) ResponseData(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {

	w.WriteHeader(status)
	o._responseData.Status = status
	o._responseData.Data = data
	ret, _err := json.Marshal(o._responseData)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func (o *ResponsePattern) ResponseError(w http.ResponseWriter, r *http.Request, status int, err interface{}) error {

	w.WriteHeader(status)
	o._responseError.Status = status
	o._responseError.Error = err
	ret, _err := json.Marshal(o._responseError)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func (o *ResponsePattern) ResponseErrorData(w http.ResponseWriter, r *http.Request, status int, err, data interface{}) error {

	w.WriteHeader(status)
	o._responseErrorData.Status = status
	o._responseErrorData.Error = err
	o._responseErrorData.Data = data
	ret, _err := json.Marshal(o._responseErrorData)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func (o *ResponsePattern) ResponseMessage(w http.ResponseWriter, r *http.Request, status int, message interface{}) error {

	w.WriteHeader(status)
	o._responseMessage.Status = status
	o._responseMessage.Message = message
	ret, _err := json.Marshal(o._responseMessage)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}

func (o *ResponsePattern) ResponseMessageData(w http.ResponseWriter, r *http.Request, status int, message, data interface{}) error {

	w.WriteHeader(status)
	o._responseMessageData.Status = status
	o._responseMessageData.Message = message
	o._responseMessageData.Data = message
	ret, _err := json.Marshal(o._responseMessageData)
	if _err != nil {
		return _err
	}
	o._write(w, r, ret)
	return nil
}
