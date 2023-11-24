package main

import (
	"log"

	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/naeem4265/product-store/handlers"
)

func main() {

	router := chi.NewRouter()

	router.Post("/signin", handlers.SignIn)
	router.Get("/signout", handlers.SignOut)

	router.Route("/brands", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetBrands)
		r.Get("/{id}", handlers.GetBrandById)
		r.Put("/{id}", handlers.PutBrand)
		r.Post("/", handlers.PostBrand)
		r.Delete("/{id}", handlers.DeleteBrand)
	})

	router.Route("/categories", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetCategories)
		r.Get("/{id}", handlers.GetCategoryById)
		r.Put("/{id}", handlers.PutCategories)
		r.Post("/", handlers.PostCategories)
		r.Delete("/{id}", handlers.DeleteCategories)
	})

	router.Route("/products", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetProducts)
		r.Get("/{id}", handlers.GetProductsById)
		r.Put("/{id}", handlers.PutProducts)
		r.Post("/", handlers.PostProducts)
		r.Delete("/{id}", handlers.DeleteProducts)
	})

	router.Route("/stocks", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetStocks)
		r.Get("/{id}", handlers.GetStockById)
		r.Put("/{id}", handlers.PutStocks)
		r.Post("/", handlers.PostStocks)
		r.Delete("/{id}", handlers.DeleteStocks)
	})

	router.Route("/suppliers", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetSuppliers)
		r.Get("/{id}", handlers.GetSupplierById)
		r.Put("/{id}", handlers.PutSuppliers)
		r.Post("/", handlers.PostSuppliers)
		r.Delete("/{id}", handlers.DeleteSuppliers)
	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the "token" cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value

		claims := &handlers.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// Token signature is invalid, return unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other error while parsing claims, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			// Token is not valid, return unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If token is valid, continue to the next handler
		next.ServeHTTP(w, r)
	})
}
