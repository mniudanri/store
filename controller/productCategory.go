package controller

import (
	"net/http"

	"github.com/mniudanri/store/config"
	"github.com/mniudanri/store/model"
)

// FindProducts godoc
// @Summary Find product and its category
// @Description Find product and its category
// @Tags Product
// @Accept */*
// @Produce json
// @Success 200 {object} response.ProductCategoryListResponse{}
// @Router /product-category [get]
func FindAllProductByItsCategory(w http.ResponseWriter, req *http.Request) {
	products, err := model.FindAllProductAndCategories()

	if err != nil {
		config.SetResponseByInterface(w, nil, http.StatusBadRequest, "Data Not Found")
		return
	}

	config.SetResponseByInterface(w, products, http.StatusOK, "")
}
