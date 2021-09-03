// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseOpenService

import (
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"goEasy/app/model/BaseOpenModel"
	"goEasy/app/model/BaseSysUserModel"
	"goEasy/app/model/BaseSysUserRoleModel"
	"goEasy/app/service/BaseSysRoleService"
	"goEasy/library/utils/cache"
	"goEasy/library/utils/captcha"
	"goEasy/library/utils/jwt"
	"time"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(baseOpenService)

type baseOpenService struct {
}

//登录逻辑
func (s *baseOpenService) Login(req *BaseOpenModel.LoginReqParams) (result interface{}, err error) {
	var userEntiy *BaseSysUserModel.Entity
	var userRoleEntiy []*BaseSysUserRoleModel.Entity
	// 校验 验证码
	if !captcha.VerifyString(req.CaptchaId, req.VerifyCode) {
		err = gerror.New("验证码错误")
		return
	}
	//判断账号密码是否正确
	user, err := BaseSysUserModel.M.Where("username=?", req.UserName).One()
	if err != nil || user == nil {
		err = gerror.New("管理用户不存在")
		return
	}
	user.Struct(&userEntiy)

	if userEntiy.Password != gmd5.MustEncryptString(req.Password+userEntiy.Salt) {
		err = gerror.New("管理用户密码不正确")
		return
	}
	//查询角色id
	userRoles, err := BaseSysUserRoleModel.M.Where("adminUserId", userEntiy.Id).All()
	if err != nil {
		gerror.New("管理员用户查询角色出错")
		return
	}
	userRoles.Structs(&userRoleEntiy)
	var rolesStr string
	for _, val := range userRoleEntiy {
		rolesStr += gconv.String(val.RoleId) + ","
	}
	//去掉多余逗号,
	rolesStr = rolesStr[0 : len(rolesStr)-1]
	//rolesStr = gstr.SubStr(rolesStr, 0, len(rolesStr)-1)

	//生成token
	exp := time.Now().Add(g.Cfg().GetDuration("jwt.expires", 1) * time.Hour).Unix()
	token, err := jwt.GenerateLoginToken(userEntiy, rolesStr, false, exp)

	expRefresh := time.Now().Add(g.Cfg().GetDuration("jwt.expires", 1) * time.Hour).Unix()
	tokenRefresh, err := jwt.GenerateLoginToken(userEntiy, rolesStr, true, expRefresh)
	if err != nil {
		err = gerror.New("生成token出错")
		return
	}
	//用户权限数组
	perms, err := BaseSysRoleService.S.GetPerms(rolesStr)
	if err != nil {
		return
	}
	permsJsonStr, err := gjson.Encode(perms)
	if err != nil {
		return
	}
	//permsJsonStr := gjson.New(perms).MustToIniString()
	//缓存cache数据
	adminAdminUserIdStr := gconv.String(userEntiy.Id)
	cache.GCacheRedis.Set("admin:"+adminAdminUserIdStr+":passwordVersion", userEntiy.PasswordV, 0)
	cache.GCacheRedis.Set("admin:"+adminAdminUserIdStr+":token", token, g.Cfg().GetDuration("jwt.expires", 1)*time.Hour)
	cache.GCacheRedis.Set("admin:"+adminAdminUserIdStr+":perms", permsJsonStr, 0)
	cache.GCacheRedis.Set("admin:"+adminAdminUserIdStr+":token:refresh", tokenRefresh, g.Cfg().GetDuration("jwt.expires", 1)*time.Hour)

	result = g.Map{
		"expire":        exp,
		"token":         token,
		"refreshExpire": expRefresh,
		"refreshToken":  tokenRefresh,
	}
	return result, nil
}

func (s *baseOpenService) RefreshToken(req *BaseOpenModel.LoginReqParams) (err error) {
	return
}
