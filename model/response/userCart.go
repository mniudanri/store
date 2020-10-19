package response

type UserCartCheckoutResponse struct {
	ProductID   int
	ProductName string
	Total       int
}

type UserCartDetail struct {
	UserCartProductDetailID int
	ProductID               int
	UserCartID              string
}

type UserProductCart struct {
	UserCartProductDetailID int
	ProductID               int
	ProductName             string
	Total                   int
}

type UserCartCategories struct {
	ProductID   int
	ProductName string
	Categories  string
}

type UserCartCategoriesListResponse struct {
	Data    UserCartCheckoutResponse
	Message string
}

type UserCartCheckoutListResponse struct {
	Data    []UserCartCheckoutResponse
	Message string
}

type UserProductCartListResponse struct {
	Data    UserProductCart
	Message string
}
