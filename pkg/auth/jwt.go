package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/venomuz/kegel-backend/pkg/logger"
	"time"
)

type TokenManager interface {
	GenerateAuthJWT(sub, role string, iss uint32) (access, refresh string, err error)
	ExtractClaims(inputToken string) (jwt.MapClaims, error)
}

func NewTokenManager(signingKey string, aud []string, log logger.Logger) *JwtHandler {
	return &JwtHandler{SignInKey: signingKey, Aud: aud, Log: log}
}

type JwtHandler struct {
	Sub       string
	Iss       uint32
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	Token     string
	SignInKey string
	Log       logger.Logger
}

// GenerateAuthJWT ...
func (hand *JwtHandler) GenerateAuthJWT(sub, role string, iss uint32) (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["iss"] = iss
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = role
	claims["aud"] = hand.Aud
	access, err = accessToken.SignedString([]byte(hand.SignInKey))
	if err != nil {
		hand.Log.Error("error generating access token", logger.Error(err))
		return
	}
	claims = refreshToken.Claims.(jwt.MapClaims)
	claims["iss"] = iss
	claims["sub"] = sub
	claims["role"] = role
	claims["aud"] = hand.Aud
	refresh, err = refreshToken.SignedString([]byte(hand.SignInKey))
	if err != nil {
		hand.Log.Error("error generating refresh token", logger.Error(err))
		return
	}
	return access, refresh, nil
}
func (hand *JwtHandler) ExtractClaims(inputToken string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(inputToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(hand.SignInKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		hand.Log.Error("invalid jwt token")
		return nil, err
	}
	return claims, nil
}
