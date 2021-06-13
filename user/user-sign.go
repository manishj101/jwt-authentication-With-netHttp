package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thegeniusgroup/jwt-authentication/data"
	"github.com/thegeniusgroup/jwt-authentication/structs"
	"github.com/thegeniusgroup/jwt-authentication/util"
	"golang.org/x/crypto/bcrypt"
)

// UserSignUp :
func UserSignUp(pathType string, h func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userDetails structs.UserClaims
		err := json.NewDecoder(r.Body).Decode(&userDetails)
		if err != nil {
			// If the structure of the body is wrong, return an HTTP error
			fmt.Println(err.Error())
			util.WebResponse(w, http.StatusBadRequest, "invaild input json")
			return
		}

		if userDetails.Email == "" || userDetails.Username == "" || userDetails.Password == "" {
			util.WebResponse(w, http.StatusBadRequest, "invaild request")
			return
		}

		// Generate Hash password
		hashPass, err := bcrypt.GenerateFromPassword([]byte(userDetails.Password), bcrypt.MinCost)
		if err != nil {
			util.WebResponse(w, http.StatusBadGateway, "Error While Hashing Password, Try Again")
			return
		}

		fmt.Println("hashPass is the user password in bcrypt for store in database for further communication", hashPass)

		userDetails.Password = ""
		userDetails.UserType = pathType

		// generating the JWT token for the user
		token, err := data.GenerateToken(userDetails)
		if err != nil {
			fmt.Println(err.Error())
			util.WebResponse(w, http.StatusBadGateway, "Unable to generate token")
			return
		}

		// Set token on cookies
		data.SetTokenCookies(w, token)

		r.Header.Set("Authorization", "bearer "+token)
		// Proceed
		h(w, r)
		return
	})
}
