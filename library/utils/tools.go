package utils

import (
	"bytes"
	"fmt"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
	"net"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gcharset"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

//字符串加密
func EncryptCBC(plainText, publicKey string) string {
	key := []byte(publicKey)
	b, e := gaes.EncryptCBC([]byte(plainText), key, key)
	if e != nil {
		g.Log().Error(e.Error())
		return ""
	}
	return gbase64.EncodeToString(b)
}

//字符串解密
func DecryptCBC(plainText, publicKey string) string {
	key := []byte(publicKey)
	plainTextByte, e := gbase64.DecodeString(plainText)
	if e != nil {
		g.Log().Error(e.Error())
		return ""
	}
	b, e := gaes.DecryptCBC(plainTextByte, key, key)
	if e != nil {
		g.Log().Error(e.Error())
		return ""
	}
	return gbase64.EncodeToString(b)
}

//服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

//获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

//获取相差时间
func GetHourDiffer(startTime, endTime string) int64 {
	var hour int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

//日期字符串转时间戳（秒）
func StrToTimestamp(dateStr string) int64 {
	tm, err := gtime.StrToTime(dateStr)
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return tm.Timestamp()
}

//时间戳转 yyyy-MM-dd HH:mm:ss
func TimeStampToDateTime(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d H:i:s")
}

//时间戳转 yyyy-MM-dd
func TimeStampToDate(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d")
}

//获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := ghttp.GetBytes(url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.GetInt("code") == 0 {
		city := json.GetString("city")
		return city
	} else {
		return ""
	}
}

//获取附件真实路径
func GetRealFilesUrl(r *ghttp.Request, path string) (realPath string, err error) {
	if gstr.ContainsI(path, "http") {
		realPath = path
		return
	}
	realPath, err = GetDomain(r)
	if err != nil {
		return
	}
	realPath = realPath + path
	return
}

//获取当前请求接口域名
func GetDomain(r *ghttp.Request) (string, error) {
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"]), nil
}

//获取附件相对路径
func GetFilesPath(fileUrl string) (path string, err error) {
	if !gstr.ContainsI(fileUrl, "http") {
		path = fileUrl
		return
	}
	pathInfo, err := gurl.ParseURL(fileUrl, 32)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return
	}
	path = gstr.TrimLeft(pathInfo["path"], "/")
	return
}

//获取reflect的方法的键
func GetReflectkeys(m map[string]reflect.Value) []string {
	l := len(m)
	keys := make([]string, 0, l)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// ========文件名相关start========

//下划线 to 小驼峰转换
func Case2SmallCamel(name string) (str string) {
	resKey := Lcfirst(Case2Camel(name))
	return resKey
}

//下划线 to 大驼峰转换
func Case2BigCamel(name string) (str string) {
	resKey := Case2Camel(name)
	return resKey
}

//大小驼峰 to 下划线
// 驼峰式写法转为下划线写法；大小驼峰都是用这个
func Camel2Case(name string) string {
	// buffer := bytes.NewBufferString("")
	buffer := bytes.NewBuffer([]byte{})
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.WriteRune('_')
			}
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

// ---工具写法--
// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

//把a=1&b=2转成 "{"a":1,"b":2}"
func Query2JsonString(origin string) (s string, err error) {
	if origin == "" {
		return "", err
	}
	arr := gstr.Split(origin, "&")
	var objMap map[string]string
	objMap = make(map[string]string, len(arr))
	for _, a := range arr {
		a1 := gstr.Split(a, "=")
		objMap[a1[0]] = a1[1]
	}
	s1, err := gjson.Encode(objMap)

	return gconv.String(s1), nil
}

// ========文件名相关end========

//代码生成-注册路由
/**
 * @Description: 加入路由
 * @param objRouterStr 加入路由字符串
 * @param codeFlag 目标标记
 * @return error
 */
func RegisterGenCodeRouter(originRouterStr string, objRouterStr string, codeFlag string) error {
	//读取文件
	routerStr := gfile.GetContents(originRouterStr)
	//判断是否已有要加入的路由字符串
	pd := gstr.Contains(routerStr, objRouterStr)
	//不存在，在加入路由
	if pd == false {
		router1 := gstr.StrTillEx(routerStr, codeFlag)
		router2 := gstr.StrEx(routerStr, codeFlag)
		//插入对应字符串
		routerStr = router1 + "\n\t\t" + objRouterStr + "\n\t\t" + codeFlag + "\n\t\t" + router2
		gfile.PutContents(originRouterStr, routerStr)
	}
	return nil
}

/**
 * @Description: 生成路由。admin api的路由
 * @param tableName 输入表名，大驼峰
 * @return error
 */
func RegisterTableNameAllRouter(tableName string, serviceGenPath string) error {
	tableNameLower := Lcfirst(tableName)
	err := RegisterGenCodeRouter(serviceGenPath+"router/adminRouter.go", `group.ALL("/`+tableNameLower+`", Admin.`+tableName+`)`, "/**insert_router_end**/")
	if err != nil {
		return gerror.New("加入admin路由出错：" + `group.ALL("/` + tableNameLower + `", Admin.` + tableName + `)`)
	}
	err = RegisterGenCodeRouter(serviceGenPath+"router/apiRouter.go", `group.ALL("/`+tableNameLower+`", Api.`+tableName+`)`, "/**api_insert_router_end**/")
	if err != nil {
		return gerror.New("加入api路由出错：" + `group.ALL("/` + tableNameLower + `", Api.` + tableName + `)`)
	}

	return err
}

//把数组变成字符串，必须数组字符类型
//  ArrToString
//  @Description: 把数组变成字符串，必须数组字符类型
// arr [a,b,c]
// sp 分割符号
//  @return string a,b,v
//
func ArrToString(arr []string, sp string) string {
	var newArr = garray.NewArrayFrom(gconv.Interfaces(arr))

	newStr := newArr.Join(sp)
	return newStr
}

//
//  StringSplitToArr
//  @Description: 字符串分割成数据
// origin "a,b,c"
// sp	分割符号
//  @return []string [a,b,c]
//
func StringSplitToArr(origin string, sp string) []string {
	arr := gstr.Split(origin, sp)
	return arr
}

//
//  checkPerms
//  @Description: 判断是否有权限
// perms 已有的权限
// uri	访问的权限
//  @return bool
//
func CheckPerms(cachePerms string, uri string) bool {
	perms, err := gjson.Decode(cachePerms)
	if err != nil {
		return false
	}
	if gstr.Contains(uri, "/admin/") {
		uri = gstr.Replace(uri, "/admin/", "")
	}

	arr := gstr.Split(uri, "/")
	uri = gstr.Join(arr, ":")
	var newArr = garray.NewArrayFrom(gconv.Interfaces(perms))

	if newArr.Contains(uri) {
		return true
	}
	return false
}
