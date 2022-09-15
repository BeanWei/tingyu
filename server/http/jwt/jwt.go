package jwt

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/BeanWei/tingyu/g"
	"github.com/BeanWei/tingyu/pkg/shared"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	*jwt.RegisteredClaims
	UserInfo *shared.CtxUser
}

// CreateToken .
func CreateToken(data *shared.CtxUser) (tokenStr string, expire time.Time, err error) {
	expire = time.Now().Add(time.Hour * 24 * time.Duration(g.Cfg().JWT.TimeoutDays))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
		data,
	})
	tokenStr, err = token.SignedString([]byte(g.Cfg().JWT.SecretKey))
	return
}

// ParseToken .
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing algorithm")
		}
		return []byte(g.Cfg().JWT.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// GetToken .
func GetToken(ctx context.Context, c *app.RequestContext) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("auth header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("auth header is invalid")
	}

	return parts[1], nil
}

// RefreshToken .
func RefreshToken(ctx context.Context, c *app.RequestContext) (tokenStr string, expire time.Time, err error) {
	token, err := GetToken(ctx, c)
	if err != nil {
		return
	}

	claims, err := ParseToken(token)
	if err != nil {
		return
	}
	if !claims.VerifyExpiresAt(time.Now(), true) {
		err = jwt.ErrTokenExpired
		return
	}
	return CreateToken(claims.UserInfo)
}

// ExtractClaims .
func ExtractClaims(ctx context.Context, c *app.RequestContext) (*Claims, error) {
	token, err := GetToken(ctx, c)
	if err != nil {
		return nil, err
	}
	claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}
	if !claims.VerifyExpiresAt(time.Now(), true) {
		return nil, jwt.ErrTokenExpired
	}
	return claims, nil
}
