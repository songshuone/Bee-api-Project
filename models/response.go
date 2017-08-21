package models
type Response struct {
	Message string `json:"message"`
	Status  int `json:"status"`
}
type ResponseData struct {
	Response
	Result interface{} `json:"result"`
}
