package model

import (
  "github.com/mniudanri/store/config"
)

type ProductCategory struct {
	ProductID      int
	ProductName    string
  Categories     string
}

func FindAllProductAndCategories() ([]ProductCategory, error) {
  productCategories := []ProductCategory{}
  conf := config.Config

  // TODO: add custom query for dinamic WHERE Clause sql?
  productCategorySql := `
    SELECT
      t1.product_id AS "productId",
    	t1.product_name AS "productName",
    	STRING_AGG(t3.category_name, ';') as "categories"
    FROM product t1
    INNER JOIN product_category t2 on t2.product_id = t1.product_id
    INNER JOIN category t3 on t3.category_id = t3.category_id
    GROUP BY t1.product_name, t1.product_id;`

	rows, err := conf.DB.Query(productCategorySql)

  if err != nil {
  	return nil, err
	}

  defer rows.Close()

	for rows.Next() {
		productCategory := ProductCategory{}

		err = rows.Scan(&productCategory.ProductID, &productCategory.ProductName, &productCategory.Categories)
		if err != nil {
			return productCategories, err
		}

		productCategories = append(productCategories, productCategory)
	}

  return productCategories, err
}
