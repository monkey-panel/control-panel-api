package codes

import "net/http"

const (
	/* http codes */
	OK           Code = http.StatusOK           // 200
	BadRequest   Code = http.StatusBadRequest   // 400
	Unauthorized Code = http.StatusUnauthorized // 401
	Forbidden    Code = http.StatusForbidden    // 403
	/* error code */
	UnknownAccount      Code = 1001
	UnknownToken        Code = 1002
	UnknownUser         Code = 1005
	UnknownInstanceCode Code = 1006

	InvalidUsername Code = 2001
	InvalidPassword Code = 2002

	UsernameAlreadyExists Code = 3001
)

type (
	Code     uint32
	CodeData struct {
		HttpCode uint16 `json:"-"`
		Code     Code   `json:"code"`
		Message  string `json:"message,omitempty"`
	}
)

type ResponseData[D any] struct {
	CodeData
	Data D `json:"data,omitempty"`
}

var codesMap = map[Code]CodeData{
	OK:         {200, OK, "OK"},
	BadRequest: {400, BadRequest, "Bad Request"},
	Forbidden:  {403, Forbidden, "Forbidden"},

	UsernameAlreadyExists: {400, UsernameAlreadyExists, "Username Already Exists"},
}

func (c Code) Base() CodeData { return codesMap[c] }
func (c Code) Error() string  { return c.Base().Message }

func Response[D any](code Code, data D) ResponseData[D] {
	return ResponseData[D]{
		CodeData: code.Base(),
		Data:     data,
	}
}

func ResponseOK[D any](data D) ResponseData[D] { return Response[D](OK, data) }
