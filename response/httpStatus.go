package response

type HttpStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *HttpStatus) Success(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    h.Code,
		"message": h.Message,
		"data":    data,
	}
}
