package middlewares

import "net/http"

func DefaultHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origins", "*")
		w.Header().Set("Access-Control-Expose-Headers", "x-total-count")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
