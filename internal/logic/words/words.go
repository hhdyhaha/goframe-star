package words

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "goframe-star/api/words/v1"
	"goframe-star/internal/dao"
	"goframe-star/internal/model/do"
	"goframe-star/internal/model/entity"
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

// 查询单词 单词分页
type ListInput struct {
	Uid  uint
	Word string
	Page int
	Size int
}

func (w *Words) List(ctx context.Context, in ListInput) (list []entity.Words, total int, err error) {
	// 对于查询初始值的处理 因为默认为0
	if in.Page == 0 {
		in.Page = 1
	}
	if in.Size == 0 {
		in.Size = 1
	}
	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)

	// 组成查询链
	if in.Uid > 0 {
		orm = orm.Where(cls.Uid, in.Uid)
	}

	// 模糊查询
	if len(in.Word) != 0 {
		orm = orm.WhereLike(cls.Word, "%"+in.Word+"%")
	}
	orm = orm.OrderDesc(cls.CreatedAt).OrderDesc(cls.Id).Page(in.Page, in.Size)
	if err = orm.ScanAndCount(&list, &total, true); err != nil {
		return
	}
	return
}
