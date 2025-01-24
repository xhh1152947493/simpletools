package middlewares

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"net/http"
	"simpletools/internal/data"
	"strconv"
	"sync"
	"time"
)

const (
	nonceExpiration = 5 * time.Minute // 时间窗口，用于防止重放攻击

	errCodeOk            = 0
	errCodeParamsInvalid = 1
	errCodeJwtInvalid    = 2
	errCodeTsInvalid     = 3
	errCodeNonceInvalid  = 4
)

type CustomClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Platform string `json:"platform"`
}

var (
	usedNonces = sync.Map{}
	jwtKey     = []byte("BFA31985-E2C7-4F6D-8A9B-73D5109E8C3A") // 不暴露给客户端，只在服务器使用
)

func GenerateToken(username, platform string) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    "simpletools-admin",
		},
		Username: username,
		Platform: platform,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 对称加密、单服验证
	return token.SignedString(jwtKey)
}

func validateJwtToken(c *gin.Context, jwtToken string) bool {
	token, err := jwt.Parse(jwtToken, func(tk *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims) // CustomClaims的jwt.StandardClaims匿名继承则实现了Valid接口，则Claims是jwt.MapClaims类型
	if !ok {
		return false
	}
	// 在后续请求中直接获取用户信息
	c.Set("username", claims["username"])
	c.Set("platform", claims["platform"])
	return true
}

// validate 由https保证请求过程中的内容不会被篡改和窃取
func validate(c *gin.Context, checkJwt bool) int {
	nonce := c.GetHeader("X-Nonce")
	timestamp := c.GetHeader("X-Timestamp")
	if nonce == "" || timestamp == "" {
		return errCodeParamsInvalid
	}
	if checkJwt {
		jwtToken := c.GetHeader("X-Authorization")
		if !validateJwtToken(c, jwtToken) {
			return errCodeJwtInvalid
		}
	}
	ts, _ := strconv.Atoi(timestamp)
	if time.Now().Unix()-int64(ts) > int64(nonceExpiration.Seconds()) {
		return errCodeTsInvalid
	}
	if _, ok := usedNonces.Load(nonce); ok {
		return errCodeNonceInvalid
	}
	usedNonces.Store(nonce, time.Now())
	return errCodeOk
}

func Validate(checkJwt bool) gin.HandlerFunc {
	// 启动一个后台协程来定期清理过期的 nonce
	go func() {
		for {
			time.Sleep(nonceExpiration)
			now := time.Now()
			usedNonces.Range(func(key, value interface{}) bool {
				if timestamp, ok := value.(time.Time); ok {
					if now.Sub(timestamp) > nonceExpiration {
						usedNonces.Delete(key)
					}
				}
				return true
			})
		}
	}()
	return func(c *gin.Context) {
		if code := validate(c, checkJwt); code != 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			data.Log().Error().HttpRequest(c.Request).Int("code", code).Msg("request sign failed")
			return
		}
		c.Next()
	}
}
