// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysUserService

import (
	"gfEasy/app/model/BaseSysUserModel"
	"gfEasy/app/model/BaseSysUserRoleModel"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/util/grand"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var S = new(baseSysUserService)

type baseSysUserService struct {
}

// 查询单个信息
func (s *baseSysUserService) Info(id int) (data *BaseSysUserModel.InfoResponse, err error) {
	var arr []*BaseSysUserRoleModel.Entity
	err = BaseSysUserModel.M.Where("id=?", id).Scan(&data)
	err = BaseSysUserRoleModel.M.Where("adminUserId", id).Scan(&arr)
	data.Password = ""
	for _, v := range arr {
		data.RoleIdList = append(data.RoleIdList, v.RoleId)
	}
	return
}

//添加操作
func (s *baseSysUserService) Add(req *BaseSysUserModel.AddReqParams) (insId int64, err error) {
	req.PasswordV = 1
	req.Salt = grand.S(6)
	req.Password = gmd5.MustEncryptString(req.Password + req.Salt)

	r, err := BaseSysUserModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()

	//角色关系绑定
	roleMenu := garray.NewArray()
	for _, v := range req.RoleIdList {
		roleMenu.Append(g.Map{
			"roleId":      v,
			"adminUserId": id,
		})
	}
	BaseSysUserRoleModel.M.OmitEmpty().Data(roleMenu).Insert()
	return id, e
}

//修改操作
func (s *baseSysUserService) Update(req *BaseSysUserModel.UpdateReqParams) (Id int, err error) {
	var user *BaseSysUserModel.Entity
	err = BaseSysUserModel.M.Where("id=?", req.Id).Scan(&user)
	if err != nil {
		return
	}
	if req.Password != "" && req.Password != user.Password && gmd5.MustEncryptString(req.Password+user.Salt) != user.Password {
		req.PasswordV += 1
		req.Password = gmd5.MustEncryptString(req.Password + user.Salt)
	} else {
		req.Password = ""
	}
	_, err = BaseSysUserModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)

	//角色关系绑定
	roleMenu := garray.NewArray()
	for _, v := range req.RoleIdList {
		roleMenu.Append(g.Map{
			"roleId":      v,
			"adminUserId": id,
		})
	}
	//先删掉
	BaseSysUserRoleModel.M.Where("adminUserId", id).Delete()
	//再绑定
	BaseSysUserRoleModel.M.OmitEmpty().Data(roleMenu).Insert()
	return id, err
}

//分页查询
func (s *baseSysUserService) Page(req *BaseSysUserModel.PageReqParams) (total, page int, size int, list []*BaseSysUserModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := BaseSysUserModel.M

	if req.Status != -1 {
		M = M.Where("status=?", req.Status)
	}

	if len(req.DepartmentIds) != 0 {
		M = M.WhereIn("departmentId", req.DepartmentIds)
	}

	if req.KeyWord != "" {
		M = M.WhereLike("name", "%"+req.KeyWord+"%")

	}
	if req.StartTime != "" {
		M = M.WhereGTE("createTime", req.StartTime)
	}
	if req.EndTime != "" {
		M = M.WhereLTE("createTime", req.EndTime)
	}
	if req.Order != "" && req.Sort != "" {
		M = M.Order(req.Order + " " + req.Sort)
	}

	total, err = M.Count()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取总行数失败")
		return
	}
	//是否是导出excel，不用分页查询
	if req.IsExport {
		exportData, err1 := M.All()
		if err1 != nil {
			g.Log().Error(err1)
			err = gerror.New("获取导出数据失败")
			return
		}
		err = exportData.Structs(&list)
		return
	}

	M = M.Page(req.Page, req.Size)

	data, err := M.All()

	if err != nil {
		g.Log().Error(err)
		err = gerror.New("获取数据失败")
		return
	}
	list = make([]*BaseSysUserModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *baseSysUserService) List(condition g.Map) (list []*BaseSysUserModel.Entity, err error) {

	err = BaseSysUserModel.M.Where(condition).Scan(&list)
	return

}

//删除
func (s *baseSysUserService) Delete(ids []int) (err error) {
	_, err = BaseSysUserModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	_, err = BaseSysUserRoleModel.M.Where("adminUserId IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除userRole模型数据失败")
		return
	}
	return
}

//部门转移
func (s *baseSysUserService) Move(req *BaseSysUserModel.MoveReqParams) (err error) {
	for _, v := range req.AdminUserIds {
		_, err = BaseSysUserModel.M.OmitEmpty().Data(
			g.Map{"departmentId": req.DepartmentId}).Where("id=?", v).Update()
		if err != nil {
			return
		}
	}

	return
}
