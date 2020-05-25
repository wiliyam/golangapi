package commonmodules

//HTTPError400 400 error model
type HTTPError400 struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"Bad request"`
}

//HTTPError401 401 error model
type HTTPError401 struct {
	Status  int    `json:"status" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

//HTTPError404 404 error model
type HTTPError404 struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"Not found"`
}

//HTTPError500 500 error model
type HTTPError500 struct {
	Status  int    `json:"status" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

// // NewError example
// func NewError(ctx *gin.Context, status int, err error) {
// 	er := HTTPError{
// 		Code:    status,
// 		Message: err.Error(),
// 	}
// 	ctx.JSON(status, er)
// }
