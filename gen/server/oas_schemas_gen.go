// Code generated by ogen, DO NOT EDIT.

package sprut_server

import (
	"fmt"
)

func (s *ErrorResponseStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Default error object.
// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorResponseStatusCode wraps Error with StatusCode.
type ErrorResponseStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorResponseStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorResponseStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorResponseStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorResponseStatusCode) SetResponse(val Error) {
	s.Response = val
}

func (*ErrorResponseStatusCode) getPingRes() {}

// Ref: #/components/schemas/Pong
type Pong struct {
	Pong bool `json:"pong"`
}

// GetPong returns the value of Pong.
func (s *Pong) GetPong() bool {
	return s.Pong
}

// SetPong sets the value of Pong.
func (s *Pong) SetPong(val bool) {
	s.Pong = val
}

func (*Pong) getPingRes() {}
