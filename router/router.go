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

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
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
