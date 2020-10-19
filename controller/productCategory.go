package controller

import (
  // "log"
  "net/http"
  // "strconv"
  "github.com/mniudanri/store/config"
  "github.com/mniudanri/store/model"
  // "github.com/go-chi/chi"
)

// get list products
func FindAllProductByItsCategory(w http.ResponseWriter, req *http.Request) {
  products, err := model.FindAllProductAndCategories()

  if err != nil {
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
    return
	}

  config.SetResponseByInterface(w, products, http.StatusOK, "")
}