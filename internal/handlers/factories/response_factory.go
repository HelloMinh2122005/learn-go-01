package factories

import (
	"minh.com/go-rest-gin-3/internal/handlers/dtos/responses"
)

func CreateSingleResponse[T any](
	statusCode int,
	message string,
	data T,
) *responses.BaseSingleResponse[T] {
	return &responses.BaseSingleResponse[T]{
		BaseResponse: responses.BaseResponse{
			Status:  statusCode,
			Message: message,
		},
		Data: data,
	}
}

func CreateListResponse[T any](
	statusCode int,
	message string,
	data []T,
) *responses.BaseListResponse[T] {
	return &responses.BaseListResponse[T]{
		BaseResponse: responses.BaseResponse{
			Status:  statusCode,
			Message: message,
		},
		Total: len(data),
		Data:  data,
	}
}
