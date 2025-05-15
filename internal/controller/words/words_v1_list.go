package words

import (
	"context"

	"goframe-star/api/words/v1"
	"goframe-star/internal/logic/words"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}

	// 调用逻辑接口,返回的是这三个参数
	wordList, total, err := c.words.List(ctx, words.ListInput{
		Uid:  uid,
		Word: req.Word,
		Page: req.Page,
		Size: req.Size,
	})

	if err != nil {
		return nil, err
	}

	// 格式化响应数据
	var list []v1.List
	for _, v := range wordList {
		list = append(list, v1.List{
			Id:               v.Id,
			Word:             v.Word,
			Definition:       v.Definition,
			ProficiencyLevel: v1.ProficiencyLevel(v.ProficiencyLevel),
		})
	}
	// 返回响应结果
	return &v1.ListRes{
		List:  list,
		Total: uint(total),
	}, nil
}
