package hello

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"

	v1 "github.com/sanrentai/gflow/api/hello/v1"
)

func init() {
	g.View().BindFuncMap(gview.FuncMap{
		"dom": Dom,
	})
}

func Dom(name string, a string, children ...string) string {
	// var a = make([]string, 0)

	// for k, v := range attrs {
	// 	a = append(a, fmt.Sprintf(` %s="%s"`, k, v))
	// }
	if a != "" {
		a = " " + a
	}
	return fmt.Sprintf(`<%s%s>%s<%s>`,
		name,
		a,
		strings.Join(children, ""),
		name,
	)
}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.WriteTpl("main.html")
	return
}
