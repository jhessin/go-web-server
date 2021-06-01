package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	// traffic
	"github.com/pilu/traffic"

	// go-guardian
	"github.com/shaj13/go-guardian/auth"
	"github.com/shaj13/go-guardian/auth/strategies/basic"
	"github.com/shaj13/go-guardian/auth/strategies/bearer"
	"github.com/shaj13/go-guardian/store"

	// mongo
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// bcrypt
	"golang.org/x/crypto/bcrypt"

	// jwt
	jwt "github.com/dgrijalva/jwt-go"
)

type any = interface{}

var authenticator auth.Authenticator

//var cache store.Cache
var cache store.Cache

func init() {
	setupGoGuardian()
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "auth-app",
			"sub": "medium",
			"aud": "any",
			"ext": time.Now().Add(time.Hour * 24).Unix(),
		})

	if jwtToken, err := token.SignedString([]byte("secret")); err == nil {
		fmt.Fprint(w, jwtToken)
	}
}

func validateUser(ctx context.Context, r *http.Request, username, pass string) (auth.Info, error) {

	fmt.Printf("Parameters\n\tUsername: %v\n\tPassword: %v", username, pass)

	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	fmt.Printf("Logging in with\nemail: %v\npassword: %v\n", email, password)

	userResult := users.FindOne(context.TODO(), bson.D{{"email", email}}, options.FindOne())

	var userBson bson.D
	if err := userResult.Decode(&userBson); err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
			return nil, err
		}
	}
	user := newUserFromBson(userBson)
	if err := bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password)); err != nil {
		//fmt.Printf("Passwords do not match: %v\n", err.Error())
		//fmt.Printf("User data: %+v\n", user)
		fmt.Printf("Error: %v\nRegistering now\n", err)
		return nil, err
	} else {
		fmt.Printf("User successfully logged in: %+v\n", user)
		return auth.NewDefaultUser(email, user.String(), nil, nil), nil
	}
}

func verifyToken(ctx context.Context, r *http.Request, tokenString string) (auth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["_id"].(string)
		var userBson bson.D
		if err := users.FindOne(ctx, bson.D{{"_id", id}}, options.FindOne()).Decode(&userBson); err != nil {
			return nil, err
		} else {
			user := newUserFromBson(userBson)
			authUser := auth.NewDefaultUser(user.email, id, nil, nil)
			return authUser, nil
		}
	}
	return nil, fmt.Errorf("invalid token")
}

func setupGoGuardian() {
	fmt.Println("Setting up Go Guardian...")
	authenticator = auth.New()
	cache = store.NewFIFO(context.Background(), time.Minute*5)
	basicStrategy := basic.New(validateUser, cache)
	tokenStrategy := bearer.New(verifyToken, cache)
	authenticator.EnableStrategy(basic.StrategyKey, basicStrategy)
	authenticator.EnableStrategy(bearer.CachedStrategyKey, tokenStrategy)
}

func authMiddleware(next http.HandlerFunc) traffic.HttpHandleFunc {
	return traffic.HttpHandleFunc(func(w traffic.ResponseWriter, r *traffic.Request) {
		log.Println("Executing Auth Middleware")
		user, err := authenticator.Authenticate(r.Request)
		if err != nil {
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			fmt.Println(err)
			return
		}
		log.Printf("User %s Authenticated\n", user.UserName())
		next(w, r.Request)
	})
}
