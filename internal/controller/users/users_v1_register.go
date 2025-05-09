package users

import (
	"context"

	"goframe-star/api/users/v1"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 调用用户逻辑层注册用户
	err = c.users.Register(ctx, req.Username, req.Password, req.Email)
	return nil, err
}
