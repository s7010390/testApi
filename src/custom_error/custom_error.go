package custom_error

type SystemError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (system SystemError) Error() string {
	return system.Message
}

type BusinessError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (business BusinessError) Error() string {
	return business.Message
}
