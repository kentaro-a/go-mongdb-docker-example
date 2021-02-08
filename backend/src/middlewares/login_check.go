package middlewares

import (
	"app/handlers"
	"app/sessions"
	"context"
	"net/http"
)

func LoginCheckMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if is_logged_in, sess := sessions.CheckSession(r); is_logged_in {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "sess", sess)
			r = r.WithContext(ctx)
			next(w, r)
		} else {
			handlers.Login(w, r)
		}
	}
}
