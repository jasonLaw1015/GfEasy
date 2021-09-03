package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"goEasy/app/model/BaseSysUserModel"
)

//
//  user 用户信息
//  roleIds 用户所属的角色id
//  isRefresh 是否是刷新token
//  exp 过期时间
//  @return string 返回token
//  @return error
//
func GenerateLoginToken(user *BaseSysUserModel.Entity, roleIds string, isRefresh bool, exp int64) (string, error) {
	//adminUserId := user.Id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"isRefresh":       isRefresh,
		"roleIds":         roleIds,
		"username":        user.Username,
		"adminUserId":     user.Id,
		"passwordVersion": user.PasswordV,
		"jwtVersion":      g.Cfg().GetString("jwt.version", "1.0"),
		"exp":             exp,
	})
	tokenString, err := token.SignedString([]byte(g.Cfg().GetString("jwt.sign", "goEasy")))
	if err == nil {
		tokenStringMd5 := gmd5.MustEncryptString(tokenString)
		g.Redis().Do("HSET", "VerifyLoginToken",
			tokenStringMd5, tokenString)
	}

	return tokenString, err
}

// 解析token
// claims["adminUserId"]这样使用
func ParseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	if tokenString == "" {
		err := gerror.New("token 为空")
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

/**
token有效正确返回用户id
*/
//func VerifyLoginToken(tokenString string) (uint, err error) {
//	//if tokenString == "" {
//	//	err = gerror.New("token不能为空")
//	//	return 0, err
//	//}
//
//}

/**
清掉所以的相关的redis
*/
func Layout(adminUserId int, tokenString string) {
	if tokenString == "" {
		return
	}
	//g.Redis().Do("HDEL", "VerifyLoginToken", gmd5.MustEncryptString(tokenString))
	//// 删除
	//g.Redis().Do("HDEL", "VerifyLoginTokenAdminUserId", adminUserId)
}
