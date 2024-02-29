package main

import (
	"context"
	"net/http"
)

type keyType string

const UserKey keyType = "req.user"

func withUser(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, UserKey, uid)
}

func getUser(ctx context.Context) string {
	return ctx.Value(UserKey).(string)
}

func (app *App) sessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, err := app.psg.AuthenticateRequest(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		ctx := withUser(r.Context(), uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
