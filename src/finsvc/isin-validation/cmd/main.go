package main

import (
	"finsvc/isin-validation/pkg"
	"finsvc/isin-validation/pkg/validation"
	"net/http"
)

func main() {
	endpoints := make(map[string]func(writer http.ResponseWriter, request *http.Request))
	endpoints["/isin/"] = pkg.IsinEndpointGenerator(validation.ValidateIsin)
	endpoints["/health/"] = pkg.HealthCheckEndpointGenerator()

	pkg.StartServer(":8080", endpoints)
}
