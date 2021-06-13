package main

import (
	"log"
	"net/http"

	user "github.com/workspace-git/jwt-authentication-With-netHttp/User"
	"github.com/workspace-git/jwt-authentication-With-netHttp/data"
	"github.com/workspace-git/jwt-authentication-With-netHttp/util"
)

var mySigningKey = []byte("mysupersecretphrase")

func SignedUpPage(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("Authorization")
	util.WebResponse(w, http.StatusOK, "You have successfully Signed Up, please use this Header: Authorization: "+r.Header.Get("Authorization"))
	return
}

func homePage(w http.ResponseWriter, r *http.Request) {
	util.WebResponse(w, http.StatusOK, "Yor are authorized user: Welcome")
	return
}

func handleRequest() {
	http.Handle("/check", data.IsAuthorized("admin", homePage))
	http.Handle("/getToken", user.UserSignUp("admin", SignedUpPage))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	handleRequest()
}
