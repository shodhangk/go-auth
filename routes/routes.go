package routes

import (
	"net/http"

	"github.com/shodhangk/go-auth/controllers"
	"github.com/shodhangk/go-auth/utils/auth"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/items", controllers.Items).Methods("GET")
	r.HandleFunc("/item/{id}", controllers.Item).Methods("GET")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	s.HandleFunc("/users", controllers.FetchUsers).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	s.HandleFunc("/orders", controllers.Orders).Methods("GET")
	s.HandleFunc("/place_order", controllers.CreateOrder).Methods("POST")

	s.HandleFunc("/cart", controllers.CartItems).Methods("GET")
	s.HandleFunc("/cart", controllers.AddToCart).Methods("POST")
	s.HandleFunc("/cart", controllers.UpdateCart).Methods("PUT")
	s.HandleFunc("/cart/:item_id", controllers.UpdateCart).Methods("DELETE")

	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
