package main

import (
	_ "github.com/sanrentai/gflow/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/sanrentai/gflow/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
