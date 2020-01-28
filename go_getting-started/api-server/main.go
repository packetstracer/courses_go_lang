package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

func main() {
	println("Starting server...")
	controllers.RegisterControllers()
	println("Server started on port 3003")
	http.ListenAndServe(":3003", nil)
}
