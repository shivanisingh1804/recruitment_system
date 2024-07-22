// package main

// import (
// 	"log"
// 	"net/http"
// 	"recruitment-system/config"
// 	"recruitment-system/routes"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	r := mux.NewRouter()
// 	routes.RegisterRoutes(r)
// 	config.ConnectDB()

//		log.Println("Server running on port 3000")
//		log.Fatal(http.ListenAndServe(":3000", r))
//	}
package main

import (
	"log"
	"net/http"
	"recruitment-system/config"
	"recruitment-system/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv() // Load environment variables
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	config.ConnectDB()

	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
