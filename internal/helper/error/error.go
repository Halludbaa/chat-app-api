package rerror


type ResponseError struct {
	Code 	int 	`json:"status,omitempty"`
	Message string 	`json:"message,omitempty"`
}

func (re *ResponseError) Error() string{
	return re.Message
}

func (re *ResponseError) GetCode() int {
	return re.Code
}

func NewResError(code int, message ...string) error{
	return &ResponseError{
		Code: code,
		Message: message[0],
	}

}

var (
	ErrNotFound = NewResError(404, "Not Found")
	ErrInternalServer = NewResError(500, "Internal Server Error")
	ErrConflict = NewResError(409, "Error Conflict")
	ErrBadReq = NewResError(400, "Bad Request Error")
	ErrUnauthorized = NewResError(401, "Unauthorized")
)