package model

import (
  "github.com/mniudanri/store/config"
  responseModel "github.com/mniudanri/store/model/response"
  "fmt"
)

func CheckoutAllProducts() (string, error) {
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
      AND t1.is_paid = FALSE
      LIMIT 1`
	_, err := conf.DB.Query(sql, 1)

  if err != nil {
  	return "Failed to checkout", err
	}

  updateSql := `
    UPDATE user_cart SET is_checkout = TRUE WHERE user_id = $1 AND is_paid = $2 AND is_checkout = $3`
	_, err = conf.DB.Exec(updateSql, 1, false, false)

  if err != nil {
  	return "Failed to checkout", err
	}
  return "Success", nil
}

func GetDetailCheckout(userCartId int) ([]responseModel.UserCartCheckoutResponse, error) {
  conf := config.Config
  checkoutModels := []responseModel.UserCartCheckoutResponse{}
  fmt.Println(userCartId)
  // TODO: add custom query for dinamic WHERE Clause sql?
  sql := `
    SELECT t3.product_id, t3.product_name, t2.total
    FROM user_cart t1
    INNER JOIN user_cart_detail_product t2 ON t2.user_cart_id = t1.user_cart_id
    INNER JOIN product t3 ON t3.product_id = t2.product_id
    WHERE
      t1.is_active = TRUE
      AND t1.user_id = $1
      AND t1.user_cart_id = $2
      AND t1.is_checkout = TRUE
      AND t1.is_paid = FALSE;`
	rows, err := conf.DB.Query(sql, 1, userCartId)

  if err != nil {
    return checkoutModels, err
  }

  defer rows.Close()

  for rows.Next() {
    checkoutModel := responseModel.UserCartCheckoutResponse{}

    err = rows.Scan(&checkoutModel.ProductID, &checkoutModel.ProductName, &checkoutModel.Total)
    if err != nil {
      return checkoutModels, err
    }

    checkoutModels = append(checkoutModels, checkoutModel)
  }

  return checkoutModels, nil
}
