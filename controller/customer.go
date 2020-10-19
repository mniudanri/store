package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mniudanri/store/config"
	"github.com/mniudanri/store/model"
	_ "github.com/mniudanri/store/model/entity"
	_ "github.com/mniudanri/store/model/response"
)

// FindProducts godoc
// @Summary Show all related product
// @Description Show all related product
// @Tags Product
// @Accept */*
// @Produce json
// @Success 200 {object} response.ProductListResponse{}
// @Router /products [get]
func FindAllProducts(w http.ResponseWriter, req *http.Request) {
	products, err := model.FindAllProduct()

	if err != nil {
		log.Print("Failed to execute query: ", err)
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
		return
	}

	config.SetResponseByInterface(w, products, http.StatusOK, "")
}

// FindProducts godoc
// @Summary Find product by requested productID
// @Description show object of product
// @Tags Product
// @Accept */*
// @Produce json
// @ID productID
// @Success 200 {object} response.ProductResponse{}
// @Router /product/{productID} [get]
func FindProductByRequestId(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "productID")
	productId, _ := strconv.Atoi(id)

	product, err := model.FindProductById(productId)

	if err != nil {
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
		return
	}

	config.SetResponseByInterface(w, product, http.StatusOK, "")
}

// CheckoutProducts godoc
// @Summary Checkout all product listed in user cart
// @Description Checkout all product listed in user cart
// @Tags User Cart
// @Accept */*
// @Produce json
// @Success 200 {object} response.DefaultResponse{}
// @Router /checkout [put]
func CheckoutProducts(w http.ResponseWriter, req *http.Request) {
	msg, err := model.CheckoutAllProducts()

	if err != nil {
		config.WriteResponseMessage(w, http.StatusBadRequest, err.Error())
		return
	}

	config.WriteResponseMessage(w, http.StatusOK, msg)
}

// FindCheckoutProduct godoc
// @Summary Find all status checkout product from user
// @Description Find all checkout product from user
// @Tags User Cart
// @Accept */*
// @Produce json
// @ID userCartId
// @Success 200 {object} response.UserCartCheckoutResponse{}
// @Router /checkout/detail/{userCartId} [get]
func GetDetailCheckout(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "userCartID")
	userCardId, _ := strconv.Atoi(id)

	detail, err := model.GetDetailCheckout(userCardId)

	if err != nil {
		config.SetResponseByInterface(w, detail, http.StatusBadRequest, err.Error())
		return
	}

	config.SetResponseByInterface(w, detail, http.StatusOK, "")
}
