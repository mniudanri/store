package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/mniudanri/store/controller"
	_ "github.com/mniudanri/store/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoute(port string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	pathJson := fmt.Sprintf("http://localhost%s/swagger/doc.json", port)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(pathJson), //The url pointing to API definition"
	))

	// PRODUCT
	r.Get("/products", controller.FindAllProducts)
	r.Get("/product/{productID}", controller.FindProductByRequestId)

	// CATEGORY
	r.Get("/product-category", controller.FindAllProductByItsCategory)

	// USER CART
	r.Post("/user-cart", controller.AddProductToCart)
	r.Get("/user-cart", controller.FindAllProductCart)

	// Checkout
	r.Put("/checkout", controller.CheckoutProducts)
	r.Get("/checkout/detail/{userCartID}", controller.GetDetailCheckout)

	http.ListenAndServe(port, r)
}
