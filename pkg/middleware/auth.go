package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitedToken := strings.Split(token, " ")

		fmt.Println(token, splitedToken)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}