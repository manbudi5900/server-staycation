package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatValidatorError(err error) []string {
	errors := []string{}
	// fmt.Println(err)
	errors = append(errors, err.Error())
	
	// for _, e := range err.(validator.ValidationErrors) {
	// 	errors = append(errors, e.Param())
	// }
	return errors
}