package middlewares

import (
	"context"
	"net/http"
	"server/utils"
	"strings"
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

		utils.PrettyDisplay("claims : ", claims)

		ctx := context.WithValue(r.Context(), "user_email", claims["email"])
		ctx = context.WithValue(ctx, "user_id", int64(claims["id"].(float64)))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
