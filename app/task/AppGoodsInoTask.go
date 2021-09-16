// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Task

import (
	"gfEasy/app/service/AppGoodsInfoService"
	"github.com/gogf/gf/frame/g"
)

func (s *schedule) GoodsTask() error {
	AppGoodsInfoService.S.Info(3)
	g.Log().Println("GoodsTask")
	return nil
}
