package main

import (
	"errors"
	_ "goframe-star/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/frame/g"

	"goframe-star/internal/cmd"
)

func main() {
	var err error

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	// 检查数据库是否能连接
	err = connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

// connDb 检查数据库是否能连接
func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败，请检查配置")
	}
	return nil
}
