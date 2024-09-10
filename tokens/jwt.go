package tokens

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UID   string
	Email string
	jwt.RegisteredClaims
}

func GenerateJWT(uid, email string) (signedToken string, err error) {
	claims := &Claims{
		UID:   uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
		},
	}

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Panicln("SECRET_KEY not an environment variable")
		return
	}
	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))

	return
}
