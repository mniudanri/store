package controller

import (
  "net/http"
  "encoding/json"
  "github.com/mniudanri/store/config"
  "github.com/mniudanri/store/model"
  payloadModel "github.com/mniudanri/store/model/payload"

)

// check and add product to user cart
func AddProductToCart(w http.ResponseWriter, req *http.Request) {
  requestBody := payloadModel.UserCartCreatePayload{}
  json.NewDecoder(req.Body).Decode(&requestBody)
  productId := requestBody.ProductId
  total := requestBody.Total

  if productId == 0 || total == 0 {
    config.WriteResponseMessage(w, http.StatusBadRequest, "ProductID and Total must be filled")
    return
  }

  userActiveCartId, err := model.FindActiveUserCartIdByProductId(productId)
  if err != nil {
    config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
    return
  } else if userActiveCartId == 0 {
    userCartId, err := model.CreateUserCart()
    if err != nil {
      config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
      return
    }

    _, err = model.CreateUserCartDetailProduct(userCartId, productId, total)
    if err != nil {
      config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
      return
    }
  } else {
    _, err = model.CreateUserCartDetailProduct(userActiveCartId, productId, total)
    if err != nil {
      config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
      return
    }
  }

  config.WriteResponseMessage(w, http.StatusOK, "Success adding product to cart!")
}

func FindAllProductCart(w http.ResponseWriter, req *http.Request) {
  products, err := model.FindAllProductCartInUser()

  if err != nil {
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, err.Error())
    return
	}

  config.SetResponseByInterface(w, products, http.StatusOK, "")
}
