package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go_app/helper"
	"net/http"
	"os"
	"strings"
	"time"
)

type JWTService interface {
	GenerateToken(UserID string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserId(ctx *gin.Context) string
}
type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "test_key",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "test_key"
	}
	return secretKey
}
func (j *jwtService) GenerateToken(UserID string) string {
	claims := jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	jwtString := strings.Split(token, "Bearer ")[1]
	return jwt.Parse(jwtString, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sihning method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserId(ctx *gin.Context) string {
	token, err := j.ValidateToken(ctx.GetHeader("Authorization"))
	if err != nil {
		res := helper.BuildErrorResponse("Unauthorized", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
