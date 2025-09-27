package responses

type BaseSingleResponse[T any] struct {
	BaseResponse
	Data T
}
