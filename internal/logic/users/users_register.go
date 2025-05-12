package users

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-star/internal/dao"
	"goframe-star/internal/model/do"
)

type RegisterInput struct {
	Username string
	Password string
	Email    string
}

// 用户注册
func (u *Users) Register(ctx context.Context, in RegisterInput) error {

	if err := u.checkUser(ctx, in.Username); err != nil {
		return err
	}

	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: in.Username,
		Password: u.encryptPassword(in.Password),
		Email:    in.Email,
	}).Insert()

	if err != nil {
		return err
	}

	return nil
}

// 校验用户名是否存在
func (u *Users) checkUser(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where("username", username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户名已存在")
	}
	return nil
}
