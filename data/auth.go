package data

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thegeniusgroup/jwt-authentication/structs"
	"github.com/thegeniusgroup/jwt-authentication/util"
)

var mySigningKey = []byte("mysupersecretphrase")

type StandardClaims struct {
	structs.UserClaims
	jwt.StandardClaims
}

// GenerateToken :
func GenerateToken(userDetails structs.UserClaims) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	StandardClaims := StandardClaims{
		userDetails,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenKey := jwt.NewWithClaims(jwt.SigningMethodHS256, StandardClaims)
	token, err := tokenKey.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err.Error())
		return token, fmt.Errorf("Unable to generate token")
	}
	return token, nil
}

func IsAuthorized(pathType string, h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userClaims StandardClaims
		if r.Header["Authorization"] != nil {
			tokenStrings := strings.Split(r.Header["Authorization"][0], " ")
			fmt.Println("token: ", tokenStrings[1])
			if strings.ToLower(tokenStrings[0]) != "bearer" || tokenStrings[1] == "" {
				util.WebResponse(w, http.StatusBadRequest, "Invaild Access")
				return
			}

			tkn, err := jwt.ParseWithClaims(tokenStrings[1], &userClaims, func(token *jwt.Token) (interface{}, error) {
				return mySigningKey, nil
			})

			if err != nil || !tkn.Valid {
				fmt.Println(err.Error())
				util.WebResponse(w, http.StatusUnauthorized, "Unauthorized access")
				return
			}
			fmt.Println(userClaims)
		} else {
			util.WebResponse(w, http.StatusUnauthorized, "Not Aauthorized")
			return
		}

		// Proceed
		h(w, r)
		return
	})
}

// SetTokenCookies :
func SetTokenCookies(w http.ResponseWriter, token string) {
	// Set token on cookies
	expirationTime := time.Now().Add(15 * time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:    "Authorization",
		Value:   token,
		Expires: expirationTime,
	})
}
