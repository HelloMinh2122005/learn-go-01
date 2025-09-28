package responses

type BaseResponse struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}
