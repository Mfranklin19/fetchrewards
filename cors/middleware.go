package cors

import (
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	/*
		Here we could apply our middleware to our webservice endpoints. I will add CORS
		middleware that would enable this to be accessed from a front end with a diffrent
		origin, however I will not be building the front end to take
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, UPDATE, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		handler.ServeHTTP(w, r)
	})
}
