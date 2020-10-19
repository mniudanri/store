package model

import (
  "github.com/mniudanri/store/config"
  "errors"
  "fmt"
)

type UserCartCategories struct {
	ProductID      int
	ProductName    string
  Categories     string
}

type UserCart struct {
  UserCartID    int
	UserID        int
	IsActive      string
}

type UserCartDetail struct {
  UserCartProductDetailID  int
	ProductID                int
	UserCartID               string
}

func FindActiveUserCartIdByProductId(productId int) (int, error){
  conf := config.Config
  userCart := UserCart{}
  userCartDetail := UserCartDetail{}

  // NOTE:
  // - find active user_cart, product
  // - find same product_id in cart, throw error if product exists
  productSql := "SELECT product_id FROM product WHERE product_id = $1"
  _ = conf.DB.QueryRow(productSql, productId).Scan(&userCartDetail.ProductID)
  fmt.Println("product_id: ", userCartDetail.ProductID, productId)

  // if userCartDetail.ProductID == 0 {
  // 	return userCart.UserCartID, errors.New("Product not found")
	// }

  cartSql := "SELECT user_cart_id FROM user_cart WHERE user_id = $1 is_active = $2"
  _ = conf.DB.QueryRow(cartSql, 1, 1).Scan(&userCart.UserCartID)

  if userCart.UserCartID != 0 {
    cartDetailSql := "SELECT user_cart_detail_product_id FROM user_cart_detail_product WHERE user_cart_id = $1 AND product_id = $2"
    _ = conf.DB.QueryRow(cartDetailSql, 1, productId).Scan(&userCartDetail.UserCartProductDetailID)

    if userCartDetail.UserCartProductDetailID != 0 {
      return userCart.UserCartID, errors.New("Product already in the cart")
    }
  }

  return userCart.UserCartID, nil
}

func CreateUserCart() (int, error){
  conf := config.Config
  userCartId := 0

  cartSql := "INSERT INTO user_cart (user_id, is_active) VALUES ($1, $2) RETURNING user_cart_id"
  err := conf.DB.QueryRow(cartSql, 1, 1).Scan(&userCartId)

  if err != nil {
  	return userCartId, err
	}

  return userCartId, nil
}

func CreateUserCartDetailProduct(userCartId int, productId int) (int, error){
  conf := config.Config
  userCartDetailId := 0

  cartSql := "INSERT INTO user_cart_detail_product (user_cart_id, product_id) VALUES ($1, $2) RETURNING user_cart_id"
  _,err := conf.DB.Exec(cartSql, userCartId, productId)
  err = conf.DB.QueryRow(cartSql, 1, productId).Scan(&userCartDetailId)

  if err != nil {
  	return userCartDetailId, err
	}

  return userCartDetailId, nil
}
