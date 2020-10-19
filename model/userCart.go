package model

import (
  "github.com/mniudanri/store/config"
  responseModel "github.com/mniudanri/store/model/response"
  "github.com/mniudanri/store/model/entity"
  "errors"
)

func FindActiveUserCartIdByProductId(productId int) (int, error){
  conf := config.Config
  userCart := entity.UserCart{}
  userCartDetail := responseModel.UserCartDetail{}

  // NOTE:
  // - find active user_cart, product
  // - find same product_id in cart, throw error if product exists
  productSql := "SELECT product_id FROM product WHERE product_id = $1"
  _ = conf.DB.QueryRow(productSql, productId).Scan(&userCartDetail.ProductID)

  if userCartDetail.ProductID == 0 {
  	return userCart.UserCartID, errors.New("Product not found")
	}

  cartSql := "SELECT user_cart_id FROM user_cart WHERE user_id = $1 AND is_active = $2"
  _ = conf.DB.QueryRow(cartSql, 1, true).Scan(&userCart.UserCartID)

  if userCart.UserCartID != 0 {
    cartDetailSql := "SELECT user_cart_detail_product_id FROM user_cart_detail_product WHERE user_cart_id = $1 AND product_id = $2"
    _ = conf.DB.QueryRow(cartDetailSql, userCart.UserCartID, productId).Scan(&userCartDetail.UserCartProductDetailID)

    if userCartDetail.UserCartProductDetailID != 0 {
      return userCart.UserCartID, errors.New("Product already in the cart")
    }
  }

  return userCart.UserCartID, nil
}

func CreateUserCart() (int, error){
  conf := config.Config
  userCartId := 0

  cartSql := "INSERT INTO user_cart (user_id, is_active, is_checkout, is_paid) VALUES ($1, $2, $3, $4) RETURNING user_cart_id"
  err := conf.DB.QueryRow(cartSql, 1, true, false, false).Scan(&userCartId)

  if err != nil {
  	return userCartId, err
	}

  return userCartId, nil
}

func CreateUserCartDetailProduct(userCartId int, productId int, total int) (int, error){
  conf := config.Config
  userCartId_1 := 0

  cartSql := "INSERT INTO user_cart_detail_product (user_cart_id, product_id, total) VALUES ($1, $2, $3) RETURNING user_cart_id"
  err := conf.DB.QueryRow(cartSql, userCartId, productId, total).Scan(&userCartId_1)

  if err != nil {
  	return userCartId_1, err
	}

  return userCartId_1, nil
}

func FindAllProductCartInUser() ([]responseModel.UserProductCart, error) {
  userProductCarts := []responseModel.UserProductCart{}
  conf := config.Config

  // TODO: add custom query for dinamic WHERE Clause sql?
  sql := `
    SELECT t3.product_id, t3.product_name, t2.total, t2.user_cart_detail_product_id
    FROM user_cart t1
    INNER JOIN user_cart_detail_product t2 ON t2.user_cart_id = t1.user_cart_id
    INNER JOIN product t3 ON t3.product_id = t2.product_id
    WHERE
      t1.is_active = TRUE
      AND t1.user_id = $1
      AND t1.is_checkout = FALSE
      AND t1.is_paid = FALSE`

	rows, err := conf.DB.Query(sql, 1)

  if err != nil {
  	return userProductCarts, err
	}

  defer rows.Close()

	for rows.Next() {
		userProductCart := responseModel.UserProductCart{}

		err = rows.Scan(&userProductCart.ProductID, &userProductCart.ProductName, &userProductCart.Total, &userProductCart.UserCartProductDetailID)
		if err != nil {
			return userProductCarts, err
		}

		userProductCarts = append(userProductCarts, userProductCart)
	}

  return userProductCarts, nil
}
