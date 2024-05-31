package middleware

import "net/http"

var (
	key   = "x-api-key"
	value = "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf"
)

func ApiKeyVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(key)
		if len(apiKey) == 0 || apiKey != value {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
