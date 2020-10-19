package response

import "github.com/mniudanri/store/model/entity"

type ProductCategoryListResponse struct {
	Data    []entity.ProductCategory
	Message string
}
