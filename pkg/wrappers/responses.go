package wrappers

type ResponseError struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

type CreatedResponse struct {
	Id      uint   `json:"id"`
	Message string `json:"message"`
}
