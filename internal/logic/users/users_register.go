package users

import (
	"context"
	"goframe-star/internal/dao"
	"goframe-star/internal/model/do"
)

func (u *Users) Register(ctx context.Context, username, password, email string) error {
	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: username,
		Password: password,
		Email:    email,
	}).Insert()

	if err != nil {
		return err
	}

	return nil
}
