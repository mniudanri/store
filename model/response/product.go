package response

import "github.com/mniudanri/store/model/entity"

type ProductListResponse struct {
	Data    []entity.Product
	Message string
}

type ProductResponse struct {
	Data    entity.Product
	Message string
}
