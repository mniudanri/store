package model

import (
  "fmt"
  "github.com/mniudanri/store/config"
)

type Product struct {
	ID       int
	ProductName    string
}

func FindAllProduct() ([]Product, error) {
  products := []Product{}
  conf := config.Config

  // TODO: add custom query for dinamic WHERE Clause sql?
  productSql := "SELECT product_id, product_name FROM product"
	rows, err := conf.DB.Query(productSql)

  if err != nil {
  	return nil, err
	}

  defer rows.Close()

	for rows.Next() {
		product := Product{}

		err = rows.Scan(&product.ID, &product.ProductName)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

  return products, err
}

func FindProductById(id int) (Product, error) {
  product := Product{}
  conf := config.Config

  fmt.Println(id)

  productSql := "SELECT product_id, product_name FROM product WHERE product_id = $1"
  err := conf.DB.QueryRow(productSql, id).Scan(&product.ID, &product.ProductName)

  if err != nil {
  	return product, err
	}

  return product, nil
}
