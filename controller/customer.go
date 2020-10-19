package controller

import (
  "log"
  "net/http"
  "strconv"
  "github.com/mniudanri/store/config"
  "github.com/mniudanri/store/model"
  "github.com/go-chi/chi"
)

// get list products
func FindAllProducts(w http.ResponseWriter, req *http.Request) {
  products, err := model.FindAllProduct()

  if err != nil {
  	log.Print("Failed to execute query: ", err)
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
    return
	}

  config.SetResponseByInterface(w, products, http.StatusOK, "")
}

// find product by id
func FindProductByRequestId(w http.ResponseWriter, req *http.Request) {
  id := chi.URLParam(req, "id")
  productId, _ := strconv.Atoi(id)

  product, err := model.FindProductById(productId)

  if err != nil {
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
    return
	}

  config.SetResponseByInterface(w, product, http.StatusOK, "")
}

func CheckoutProducts(w http.ResponseWriter, req *http.Request) {
  msg, err := model.CheckoutAllProducts()

  if err != nil {
		config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
    return
	}

  config.WriteResponseMessage(w, http.StatusOK, msg)
}

func GetDetailCheckout(w http.ResponseWriter, req *http.Request) {
  id := chi.URLParam(req, "UserCartId")
  userCardId, _ := strconv.Atoi(id)

  detail, err := model.GetDetailCheckout(userCardId)

  if err != nil {
		config.SetResponseByInterface(w, detail, http.StatusBadRequest, err.Error())
    return
	}

  config.SetResponseByInterface(w, detail, http.StatusOK, "")
}
