package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"goframe-star/internal/consts"
	"net/http"
)

func Auth(r *ghttp.Request) {
	// 获取请求头中的Authorization字段
	var tokenString = r.Header.Get("Authorization")

	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	// 如果解析失败,或者token无效,返回403 Forbidden
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}

	// 放在函数最后面,表示这是一个前置中间件,在请求处理之前执行
	r.Middleware.Next()
}
