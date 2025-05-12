package users

import (
	"context"

	"goframe-star/api/users/v1"
	"goframe-star/internal/logic/users"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 调用用户逻辑层注册用户
	err = c.users.Register(ctx, users.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	return nil, err
}
