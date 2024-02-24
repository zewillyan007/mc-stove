package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func Reqid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if len(r.Header.Get("reqid")) == 0 {
			r.Header.Set("reqid", uuid.New().String())
		}
		w.Header().Set("reqid", r.Header.Get("reqid"))
		next.ServeHTTP(w, r)
	})
}
