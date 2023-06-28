package handlers

type ResponseWithData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
