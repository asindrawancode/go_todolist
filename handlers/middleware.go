package handlers

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go_todolist/models"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := models.DB.Where("username = ?", claims.Subject).First(&user).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
