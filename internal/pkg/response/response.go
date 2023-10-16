package response

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type (
	Response[T any] struct {
		Code    int  `json:"-"`
		Success bool `json:"success"`
		Data    T    `json:"data"`
		Error   any  `json:"error"`
	}

	ErrorInternal struct {
		code    int
		message string
		detail  string
		source  string
		ctx     context.Context
	}

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Detail  string `json:"detail"`
		Source  string `json:"source"`
	}

	Wrapper[T any] struct{}
)

func New[T any]() *Wrapper[T] {
	return &Wrapper[T]{}
}

func (w Wrapper[T]) Send(data T, err error) (code int, response Response[T]) {
	if err != nil {
		var errWrapper *ErrorInternal

		ok := errors.As(err, &errWrapper)
		if ok {
			response = Response[T]{
				Code:    errWrapper.code,
				Success: false,
				Data:    data,
				Error: Error{
					Code:    errWrapper.code,
					Message: errWrapper.message,
					Detail:  errWrapper.detail,
					Source:  errWrapper.source,
				},
			}
		} else {
			response = Response[T]{
				Code:    http.StatusInternalServerError,
				Success: false,
				Data:    data,
				Error: Error{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
					Detail:  fmt.Sprintf("%v", err),
					Source:  getSource(runtime.Caller(1)),
				},
			}
		}
	} else {
		response = Response[T]{
			Code:    http.StatusOK,
			Success: true,
			Data:    data,
			Error:   struct{}{},
		}
	}

	return response.Code, response
}

func (w Wrapper[T]) Error() (err *ErrorInternal) {
	err = &ErrorInternal{}
	err.source = getSource(runtime.Caller(1))
	return
}

func (e *ErrorInternal) Error() string {
	return e.detail
}

func (e *ErrorInternal) SetCode(code int) *ErrorInternal {
	e.code = code
	return e
}

func (e *ErrorInternal) SetMessage(message string) *ErrorInternal {
	e.message = message
	return e
}

func (e *ErrorInternal) SetDetail(err error) *ErrorInternal {
	var data *ErrorInternal

	if errors.As(err, &data) {
		e.detail = data.detail
	} else {
		e.detail = err.Error()
	}

	return e
}

func (e *ErrorInternal) SetContext(ctx context.Context) *ErrorInternal {
	e.ctx = ctx
	return e
}

func getSource(pc uintptr, file string, line int, ok bool) (source string) {
	var funcName string

	if details := runtime.FuncForPC(pc); details != nil {
		titles := strings.Split(details.Name(), ".")
		funcName = fmt.Sprintf("%s", titles[len(titles)-1])
	}

	if ok {
		source = fmt.Sprintf("Called from %s, line #%d, func: %v", file, line, funcName)
	}

	return
}
