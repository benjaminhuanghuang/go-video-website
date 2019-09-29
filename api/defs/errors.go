package defs

type Error struct {
	Message string `json:"message"`
	// System error code , NOT http status code
	Code string `json:"code"`
}

type ErrorResponse struct {
	// Http status code
	HttpSC int
	Error  Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC: 400,
		Error: Error{
			Message: "Request body is NOT correct ",
			Code:    "001",
		},
	}

	ErrorNotAuthUser = ErrorResponse{
		HttpSC: 401,
		Error: Error{
			Message: "User authentication failed ",
			Code:    "002",
		},
	}
)
