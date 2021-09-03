// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package Task

import (
	"github.com/gogf/gf/frame/g"
	"goEasy/app/service/AppGoodsInfoService"
)

func (s *schedule) GoodsTask() error {
	AppGoodsInfoService.S.Info(3)
	g.Log().Println("GoodsTask")
	return nil
}
