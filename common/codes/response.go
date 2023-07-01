package codes

import "net/http"

const (
	/* http codes */
	OK           Code = http.StatusOK           // 200
	BadRequest   Code = http.StatusBadRequest   // 400
	Unauthorized Code = http.StatusUnauthorized // 401
	Forbidden    Code = http.StatusForbidden    // 403
	/* error code */
	UnknownAccount  Code = 1001
	UnknownToken    Code = 1002
	UnknownUser     Code = 1005
	UnknownInstance Code = 1006

	InvalidFormBody Code = 2001
	InvalidUsername Code = 2002
	InvalidPassword Code = 2003

	UsernameAlreadyExists Code = 3001
)

type (
	Code     uint32
	CodeData struct {
		HttpCode uint16 `json:"-"`
		Code     *Code  `json:"code"`
		Message  string `json:"message,omitempty"`
	}
)

type ResponseData[D any] struct {
	CodeData
	Data   D                 `json:"data,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}

var codesMap = map[Code]CodeData{
	OK:         {200, nil, "OK"},
	BadRequest: {400, nil, "Bad Request"},
	Forbidden:  {403, nil, "Forbidden"},

	/* error code */
	UnknownAccount:  {400, nil, "Unknown Account"},
	UnknownToken:    {400, nil, "Unknown Token"},
	UnknownUser:     {400, nil, "Unknown User"},
	UnknownInstance: {400, nil, "Unknown Instance"},

	InvalidFormBody: {400, nil, "Invalid Form Body"},
	InvalidUsername: {400, nil, "Invalid User Name"},
	InvalidPassword: {400, nil, "Invalid Password"},

	UsernameAlreadyExists: {400, nil, "Username Already Exists"},
}

func (c Code) Base() (d CodeData) {
	var ok bool
	if d, ok = codesMap[c]; !ok {
		d = codesMap[BadRequest]
	}
	if d.Code == nil {
		d.Code = &c
	}
	return
}
func (c Code) Error() string { return c.Base().Message }

func Response[D any](code Code, data D, errors map[string]string) (int, ResponseData[D]) {
	httpCode := int(code.Base().HttpCode)
	if errors != nil {
		return httpCode, ResponseData[D]{
			CodeData: code.Base(),
			Data:     data,
			Errors:   errors,
		}
	}
	return httpCode, ResponseData[D]{
		CodeData: code.Base(),
		Data:     data,
	}
}
