package API

import (
	"fmt"
	"log"
	"time"

	//"user-athentication-golang/database"

	jwt "github.com/dgrijalva/jwt-go"
)

// A SignedDetails holds an email and it's encoded counterpart using JWT
type SignedDetails struct {
	Email   string
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}

//var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// GenerateAllTokens generates both the detailed token and refresh token
func GenerateAllTokens(email string, name string, isAdmin bool) (signedToken string, signedRefreshToken string, err error) {
	//fmt.Println(email)
	var SECRET_KEY string = GetSecretKey()
	tokenDuration := GetTokenDuration()
	refreshTokenDuration := GetRefreshTokenDuration()
	claims := &SignedDetails{
		Email:   email,
		Name:    name,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(tokenDuration)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(refreshTokenDuration)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

// ValidateToken validates the jwt token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	var SECRET_KEY string = GetSecretKey()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		return
	}

	return claims, msg
}

// UpdateAllTokens renews the user tokens when they login
func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	return
}
