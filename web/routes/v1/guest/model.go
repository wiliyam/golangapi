package guest

type Guest struct {
	Err     string `json:"err" example:""`
	Message string `json:"message" example:"some message"`
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
