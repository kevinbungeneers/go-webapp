package middleware

import (
	"go-webapp/internal"
	"net/http"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: UUID? Something else?
		fancyUniqueId := "i_am_unique"

		ctx := r.Context()
		ctx = internal.NewRequestIdContext(ctx, fancyUniqueId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
