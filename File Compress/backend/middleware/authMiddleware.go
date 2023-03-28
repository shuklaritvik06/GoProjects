package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/shuklaritvik06/GoProjects/filecompress/backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		if utils.ValidateToken(token) == false {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		next.ServeHTTP(w, r)
	})
}
