package middlewares

import (
	"RestAPIGolang/internal/auth"
	"RestAPIGolang/internal/database"
	"RestAPIGolang/internal/helpers"
	"RestAPIGolang/internal/utils"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("x-jwt-token")
		token, err := auth.ValidateJWT(tokenString)
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}
		if !token.Valid {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}
		userID, err := helpers.GetID(r)
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}
		account, err := database.GetAccountByID(userID)
		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if account.Number != int64(claims["accountNumber"].(float64)) {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "permission denied"})
			return
		}

		if err != nil {
			utils.WriteJSON(w, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}

		handlerFunc(w, r)
	}
}
