package pkg

import (
	"encoding/json"
	"finsvc/isin-validation/pkg/validation"
	"log"
	"net/http"
	"time"
)

const ResponseOk = "{\"staus\": \"OK\"}"

func StartServer(addr string, endpoints map[string]func(writer http.ResponseWriter, request *http.Request)) {
	isinMux := http.NewServeMux()

	log.Println("Will create endpoints: ", endpoints)
	for endpoint, handler := range endpoints {
		isinMux.HandleFunc(endpoint, handler)
	}

	restHandler := accessLogMiddleware(isinMux)
	restHandler = contentTypeMiddleware(restHandler)

	log.Println("Starting server at the address: ", addr)
	err := http.ListenAndServe(addr, restHandler)
	if err != nil {
		log.Fatal("Error happened during starting server ", err)
	}
}

func contentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
}

func accessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("accessLog", r.URL.Path)
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s, %s %s\n", r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
	})
}

func IsinEndpointGenerator(validate func(isin string) *validation.Result) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		isinToCheck := r.URL.Query().Get("check")
		validationResult := *validate(isinToCheck)
		marshalledResult, err := json.Marshal(validationResult)
		if err != nil {
			log.Println("Error during marshalling of validationResult", validationResult, err)
			http.Error(w, "Internal server error", 500)
			return
		}

		_, err = w.Write(marshalledResult)
		if err != nil {
			log.Println("Error during sending the result", marshalledResult, err)
			return
		}
	}
}

func HealthCheckEndpointGenerator() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(ResponseOk))
		if err != nil {
			log.Println("Error during sending the result", err)
			return
		}
	}
}
