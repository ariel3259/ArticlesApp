package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddlewares(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	for _, r := range m {
		wrapped = r(h)
	}
	return wrapped
}
