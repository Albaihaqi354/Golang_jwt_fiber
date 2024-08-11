package helpers

import (
	"golang_jwt_copy/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mySignInKey = []byte("secretkey")

type MyCutomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.User) (string, error) {
	claims := MyCutomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySignInKey)

	return ss, err
}

func ValidationToken(tokenString string) (any, error) {
	claims := &MyCutomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySignInKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}

// func main() {
// 	user := &models.User{
// 		ID:    1,
// 		Name:  "John Doe",
// 		Email: "john@example.com",
// 	}

// 	tokenString, err := CreateToken(user)
// 	if err != nil {
// 		fmt.Println("Error creating token:", err)
// 		return
// 	}

// 	fmt.Println("Token:", tokenString)

// 	claims, err := ValidationToken(tokenString)
// 	if err != nil {
// 		fmt.Println("Error validating token:", err)
// 		return
// 	}

// 	fmt.Println("Claims:", claims)
// }
