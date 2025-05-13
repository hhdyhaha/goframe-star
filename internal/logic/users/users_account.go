package users

import (
	"context"
	"goframe-star/internal/consts"
	"goframe-star/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/golang-jwt/jwt/v5"

	"goframe-star/internal/dao"
)

type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

// 登录
func (u *Users) Login(ctx context.Context, username, password string) (tokenstring string, err error) {
	var user entity.Users
	// Scan 将查询结果赋值给 user
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}

	// 如果用户不存在,user结构体的字段保持默认值 0
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}

	// 将密码加密后与数据库中的密码进行比较
	if user.Password != u.encryptPassword(password) {
		return "", gerror.New("密码错误")
	}

	// 生成token
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	return token.SignedString([]byte(consts.JwtKey))
}
