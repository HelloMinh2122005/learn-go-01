package responses

type BaseListResponse[T any] struct {
	BaseResponse
	Total int
	Data  []T
}
