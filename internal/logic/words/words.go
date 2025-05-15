package words

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "goframe-star/api/words/v1"
	"goframe-star/internal/dao"
	"goframe-star/internal/model/do"
)

// 定义Words对象
type Words struct {
}

// 新建New函数 实例化对象
func New() *Words {
	return &Words{}
}

// 创建单词逻辑
type CreateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

// 新增单词
func (w *Words) Create(ctx context.Context, in CreateInput) error {
	var cls = dao.Words.Columns()
	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Word, in.Word).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

// 编辑单词
type UpdateInput struct {
	Uid                uint
	Word               string
	Definition         string
	ExampleSentence    string
	ChineseTranslation string
	Pronunciation      string
	ProficiencyLevel   v1.ProficiencyLevel
}

// 编辑单词
func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {
	var cls = dao.Words.Columns()
	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Word, in.Word).WhereNot(cls.Id, id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}

	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where(cls.Id, id).Where(cls.Uid, in.Uid).Update()
	if err != nil {
		return err
	}
	return nil
}
