package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"server/utils"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		tokenArr := strings.Split(bearer, " ")
		if len(tokenArr) != 2 {
			http.Error(w, "token malformed", http.StatusBadRequest)
			return
		}

		token := tokenArr[1]
		claims, err := utils.ParseToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dates := claims["date"].(map[string]interface{})
		expiresAt := int64(dates["expiresAt"].(float64))
		fmt.Println(expiresAt)

		durationSiceExpire := time.Since(time.Unix(expiresAt, 0))
		durationSec := durationSiceExpire.Seconds()
		if durationSec > 0 {
			http.Error(w, "token expired", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user_email", claims["email"])
		ctx = context.WithValue(ctx, "user_id", int64(claims["id"].(float64)))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
