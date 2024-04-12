package mid

import (
	"context"
	"net/http"
	"server/utils/encrypt"
	"strings"
)

func IsAuth(next http.Handler, isSocket bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuthenticated := true

		if !isAuthenticated {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		var token string
		if isSocket {
			token = r.URL.Query().Get("token")

		} else {

			auth_header := r.Header.Get("Authorization")
			if len(auth_header) == 0 || !strings.HasPrefix(auth_header, "Bearer") {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			token = strings.Split(auth_header, " ")[1]
		}
		claim, err := encrypt.GetClaimsFromToken(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_email", claim["Email"])
		ctx = context.WithValue(ctx, "user_id", claim["ID"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
