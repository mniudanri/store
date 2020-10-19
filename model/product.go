package model

import (
  "github.com/mniudanri/store/config"
  "github.com/mniudanri/store/model/entity"
)

func FindAllProduct() ([]entity.Product, error) {
  products := []entity.Product{}
  conf := config.Config

  // TODO: add custom query for dinamic WHERE Clause sql?
  productSql := "SELECT product_id, product_name FROM product"
	rows, err := conf.DB.Query(productSql)

  if err != nil {
  	return products, err
	}

  defer rows.Close()

	for rows.Next() {
		product := entity.Product{}

		err = rows.Scan(&product.ProductID, &product.ProductName)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

  return products, nil
}

func FindProductById(id int) (entity.Product, error) {
  product := entity.Product{}
  conf := config.Config

  productSql := "SELECT product_id, product_name FROM product WHERE product_id = $1"
  err := conf.DB.QueryRow(productSql, id).Scan(&product.ProductID, &product.ProductName)

  if err != nil {
  	return product, err
	}

  return product, nil
}
