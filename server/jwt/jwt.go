package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v2"
)

var secretKey = []byte("") // Замените на свой секретный ключ

func init(){

	filename := "./jwt/conf.yaml"
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	// Распаковка данных YAML в структуру
	type Config struct {
		Token_user string `yaml:"token_user"`
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshaling data: %v", err)
	}
	secretKey = []byte(config.Token_user)
}

func GenerateJWT(id int,login string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["login"] = login
	claims["id"] = id
	return token.SignedString(secretKey)
}

func CheckTokenOnValid(token string) bool{
	if token == "" {
		return false
	}
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false
	}

	if parsedToken.Valid {
		return true
	} else {
		return false
	}
}

func GetClaims(tokenString string) jwt.MapClaims{
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return []byte(secretKey), nil
    })
	if err != nil || !token.Valid {
        log.Printf("Error parsing token: %v", err)
		return nil
    }

	claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        log.Printf("Error getting claims from token")
        //panic("Invalid token claims")
		return nil
    }

	return claims
}