package words

import (
	"context"

	"goframe-star/api/words/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	// 调用逻辑接口
	word, err := c.words.Detail(ctx, uid, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DetailRes{
		Id:                 word.Id,
		Word:               word.Word,
		Definition:         word.Definition,
		ExampleSentence:    word.ExampleSentence,
		ChineseTranslation: word.ChineseTranslation,
		Pronunciation:      word.Pronunciation,
		ProficiencyLevel:   v1.ProficiencyLevel(word.ProficiencyLevel),
		CreatedAt:          word.CreatedAt,
		UpdatedAt:          word.UpdatedAt,
	}, nil
}
