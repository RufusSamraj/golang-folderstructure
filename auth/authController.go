package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"main.go/constant"
	"main.go/dto"
	"net/http"
	"time"
)

type AuthController struct {
}

var jwtSecret = constant.JwtSecret

// expires in 15 minutes
func GenerateAccessToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	return token.SignedString(jwtSecret)
}

// expires in 7 days
func GenerateRefreshToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

func ValidateRefreshToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

func (a AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	//var reqBody dto.RefreshTokenRequest
	//
	//// Parse the request body to get the refresh token
	//if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
	//	http.Error(w, "Invalid request body", http.StatusBadRequest)
	//	return
	//}
	//
	//// Validate the refresh token
	//refreshToken, err := ValidateRefreshToken(reqBody.RefreshToken)
	//if err != nil || !refreshToken.Valid {
	//	http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
	//	return
	//}
	//

	//claims, ok := refreshToken.Claims.(jwt.MapClaims)
	//if !ok || !refreshToken.Valid {
	//	http.Error(w, "Invalid refresh token claims", http.StatusUnauthorized)
	//	return
	//}
	//
	//rollNo, ok := claims["rollNo"].(float64)
	//if !ok {
	//	http.Error(w, "Invalid roll number in token", http.StatusUnauthorized)
	//	return
	//}

	// Generate a new access token
	newAccessToken, err := GenerateAccessToken()
	if err != nil {
		constant.WriteResponse(w, 400, err)
		return
	}

	newRefreshToken, err := GenerateRefreshToken()
	if err != nil {
		constant.WriteResponse(w, 400, err)
		return
	}

	var response dto.AuthTokenResponse

	response.AccessToken = newAccessToken
	response.RefreshToken = newRefreshToken

	constant.WriteResponse(w, 200, response)
}
