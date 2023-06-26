package common

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
		HttpCode uint16
		Code     uint   `json:"code"`
		Message  string `json:"message"`
	}
)

type ResponseData[D any] struct {
	CodeData
	Data D `json:"data"`
}

var codesMap = map[Code]CodeData{
	OK:         {200, 0, "OK"},
	BadRequest: {400, 1, "Bad Request"},
	Forbidden:  {403, 2, "Forbidden"},
}

func (c Code) Base() CodeData { return codesMap[c] }

func Response[D any](code Code, data D) ResponseData[D] {
	return ResponseData[D]{
		CodeData: code.Base(),
		Data:     data,
	}
}

func ResponseOK[D any](data D) ResponseData[D] { return Response[D](OK, data) }
