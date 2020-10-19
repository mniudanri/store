package router

import (
  "github.com/mniudanri/store/controller"
  "fmt"
  "net/http"
  "github.com/go-chi/chi"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

func InitRoute(port string) {
  r := chi.NewRouter()

  pathJson := fmt.Sprintf("http://localhost%s/swagger/doc.json", port)

  r.Get("/swagger/*", httpSwagger.Handler(
    httpSwagger.URL(pathJson), //The url pointing to API definition"
  ))

  r.Get("/products", controller.FindAllProducts)
  r.Get("/product/{id}", controller.FindProductByRequestId)
  r.Get("/product-category", controller.FindAllProductByItsCategory)
  r.Post("/user-cart", controller.AddProductToCart)

  http.ListenAndServe(port, r)
}
