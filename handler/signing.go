package handler

import (
	"auth/dto"
	"auth/model"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Create the JWT Key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Create a struct that will be encoded to a JWT
// Add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

// Signing handler
func Signing(w http.ResponseWriter, r *http.Request) {
	var cred Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our user
	// If a password exists for the given user AND if it is same as the password we received,
	// then we can move ahead. If NOT, then we return an "Unauthorized" status
	user, err := dto.Auth(cred.Username, cred.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user.Collection, _ = model.GetCollectionFromUser(user.ID)

	// Declare the expiration time for the token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT Claims, which includes the username and expiry time
	claims := &Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there's an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}
