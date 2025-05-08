package main

import (
	_ "goframe-star/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-star/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
