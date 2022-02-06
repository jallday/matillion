package models

import (
	"encoding/json"
	"fmt"
)

// swagger:model
type Error struct {
	// api error id code
	ID string `json:"id"`
	// detailed error message
	DetailedError string `json:"detailed_error"`
	// where the error occured
	Where string `json:"-"`
	// parameters relating to the error
	Params map[string]interface{} `json:"params,omitempty"`
	// the status code the relates to the error
	StatusCode int `json:"-"`
	// the id of the request
	RequestId string `json:"request_id"`
}

func NewError(where, id, detailedError string, params map[string]interface{}, code int) *Error {
	return &Error{
		Where:         where,
		ID:            id,
		DetailedError: detailedError,
		Params:        params,
		StatusCode:    code,
	}
}

func (e *Error) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// this is an important function that allows the struct to make the
// error interface and hence be passed as an error type.
func (e *Error) Error() string {
	return fmt.Sprintf("where: %s id: %s message: %s", e.Where, e.ID, e.DetailedError)
}

type ErrNotFound struct {
	Resource string
	Param    string
}

func NewErrNotFound(resource, param string) *ErrNotFound {
	return &ErrNotFound{
		Resource: resource,
		Param:    param,
	}
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("resource: %s param: %s", e.Resource, e.Param)
}

type ErrConflict struct {
	Resource string
	Param    string
}

func NewErrConflict(resource, param string) *ErrNotFound {
	return &ErrNotFound{
		Resource: resource,
		Param:    param,
	}
}

func (e *ErrConflict) Error() string {
	return fmt.Sprintf("resource: %s param: %s", e.Resource, e.Param)
}
