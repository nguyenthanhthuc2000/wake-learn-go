package response

const (
	Success               = 2000
	ErrCodeNotFound       = 4004
	ErrCodeInvalidRequest = 4005
)

// message
var msg = map[int]string{
	Success:               "Success",
	ErrCodeNotFound:       "Object not found",
	ErrCodeInvalidRequest: "Invalid request",
}
