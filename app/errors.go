package app

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ErrorCode uint
type ErrorTranslation string

const (
	ERR_EN ErrorTranslation = "EN"
	ERR_NL ErrorTranslation = "NL"
)

type Errorer interface {
	error
	ErrCode() ErrorCode
	ToJSON() map[string]interface{}
	ToJSONWithDepth(uint) map[string]interface{}
	UnWrap() []error
	Wrap([]error) Errorer
	AddToWrapped([]error) Errorer
	SetTranslation(ErrorTranslation) Errorer
}

type Error struct {
	errCode     ErrorCode
	errs        []error
	message     map[ErrorTranslation]string
	serviceName string
	translation ErrorTranslation
}

func (e Error) Error() string {
	var msgs []string
	var msgErr string
	for _, er := range e.errs {
		var err Error
		if errors.As(er, &err) {
			msgs = append(msgs, err.SetTranslation(e.translation).Error())
		} else if er != nil {
			msgs = append(msgs, er.Error())
		}
	}
	if e.translation == "" {
		msgErr = e.message[ERR_EN]
	} else {
		msgErr = e.message[e.translation]
	}
	msg := strings.Join([]string{"[", strings.Join(msgs, ", "), "]"}, "")
	return fmt.Sprintf("%v :: %s: %v", e.ErrCode(), msgErr, msg)
}

func (e Error) ErrCode() ErrorCode {
	if e.errCode != 0 {
		return e.errCode
	} else {
		for _, er := range e.errs {
			var err Error
			if errors.As(er, &err) {
				return err.ErrCode()
			}
		}
	}
	return ErrorCode(0)
}

func (e Error) UnWrap() []error {
	return e.errs
}

func (e Error) Wrap(errs []error) Errorer {
	return Error{
		errs:        errs,
		errCode:     e.errCode,
		message:     e.message,
		serviceName: e.serviceName,
	}
}

func (e Error) AddToWrapped(errs []error) Errorer {
	return Error{
		errs:        append(e.errs, errs...),
		errCode:     e.errCode,
		message:     e.message,
		serviceName: e.serviceName,
	}
}

func (e Error) SetTranslation(errorTranslation ErrorTranslation) Errorer {
	return Error{
		errs:        e.errs,
		errCode:     e.errCode,
		message:     e.message,
		translation: errorTranslation,
		serviceName: e.serviceName,
	}
}

func (e Error) ToJSON() map[string]interface{} {
	body := make(map[string]interface{}, 0)
	if e.errCode != 0 {
		body["error_code"] = strconv.Itoa(int(e.errCode))
	}
	if e.serviceName != "" {
		body["service_name"] = e.serviceName
	}
	var msgErr string
	if e.translation == "" {
		msgErr = e.message[ERR_EN]
	} else {
		msgErr = e.message[e.translation]
	}
	body["message"] = msgErr
	var jsonErrors []interface{}
	for _, er := range e.errs {
		var err Error
		if errors.As(er, &err) {
			jsonErrors = append(jsonErrors, err.SetTranslation(e.translation).ToJSON())
		} else if er != nil {
			jsonErrors = append(jsonErrors, er.Error())
		}
	}

	if len(jsonErrors) > 0 {
		body["other_errors"] = jsonErrors
	}
	return body
}

func (e Error) ToJSONWithDepth(depth uint) map[string]interface{} {
	body := make(map[string]interface{}, 0)
	if e.errCode != 0 {
		body["error_code"] = strconv.Itoa(int(e.errCode))
	}
	if e.serviceName != "" {
		body["service_name"] = e.serviceName
	}
	var msgErr string
	if e.translation == "" {
		msgErr = e.message[ERR_EN]
	} else {
		msgErr = e.message[e.translation]
	}
	body["message"] = msgErr

	var jsonErrors []interface{}
	for _, er := range e.errs {
		var err Error
		if errors.As(er, &err) {
			if depth > 0 {
				jsonErrors = append(jsonErrors, err.SetTranslation(e.translation).ToJSONWithDepth(depth-1))
			}
		} else if er != nil {
			jsonErrors = append(jsonErrors, er.Error())
		}
	}

	if len(jsonErrors) > 0 {
		body["other_errors"] = jsonErrors
	}
	return body
}

func NewErrorTranslated(message map[ErrorTranslation]string, code ErrorCode, serviceName string) Error {
	return Error{
		errCode:     code,
		message:     message,
		serviceName: serviceName,
	}
}

func NewError(message string, code ErrorCode, serviceName string) Error {
	return Error{
		errCode:     code,
		serviceName: serviceName,
		message: map[ErrorTranslation]string{
			ERR_EN: message,
		},
	}
}
