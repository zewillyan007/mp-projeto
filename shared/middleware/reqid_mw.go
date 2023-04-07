package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

func Reqid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("reqid")) == 0 {
			reqid := uuid.New()
			r.Header.Set("reqid", reqid.String())
		}
		next.ServeHTTP(w, r)
	})
}
