package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/lamhai1401/hackernews/internal/user"

	"github.com/lamhai1401/hackernews/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			users := user.User{Username: username}
			id, err := user.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			users.ID = strconv.Itoa(id)
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &users)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *user.User {
	raw, _ := ctx.Value(userCtxKey).(*user.User)
	return raw
}
