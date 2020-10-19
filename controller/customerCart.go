package controller

import (
  // "log"
  "net/http"
  // "strconv"
  "encoding/json"
  "github.com/mniudanri/store/config"
  "github.com/mniudanri/store/model"
  payloadModel "github.com/mniudanri/store/model/payload"
  // "github.com/go-chi/chi"

)

// check and add product to user cart
func AddProductToCart(w http.ResponseWriter, req *http.Request) {
  requestBody := payloadModel.UserCartCreatePayload{}
  json.NewDecoder(req.Body).Decode(&requestBody)
  productId := requestBody.ProductId

  if productId == 0 {
    config.WriteResponseMessage(w, http.StatusBadRequest, "ProductID must be filled")
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

    _, err = model.CreateUserCartDetailProduct(userCartId, productId)
    if err != nil {
      config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
      return
    }
  } else {
    _, err = model.CreateUserCartDetailProduct(userActiveCartId, productId)
    if err != nil {
      config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
      return
    }
  }

  config.WriteResponseMessage(w, http.StatusOK, "Success adding product to cart!")
}
