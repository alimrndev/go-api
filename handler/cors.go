// Package handler berisi fungsi-fungsi handler untuk aplikasi
package handler

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSMiddleware mengembalikan handler CORS yang sudah diatur
func CORSMiddleware(next http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(next)
}
