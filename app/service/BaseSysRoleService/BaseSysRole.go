// ==========================================================================
// 生成人：JasonLaw
// ==========================================================================
package BaseSysRoleService

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"goEasy/app/model/BaseSysMenuModel"
	"goEasy/app/model/BaseSysRoleDepartmentModel"
	"goEasy/app/model/BaseSysRoleMenuModel"
	"goEasy/app/model/BaseSysRoleModel"
)

var S = new(baseSysRoleService)

type baseSysRoleService struct {
}

// 查询单个信息
func (s *baseSysRoleService) Info(id int) (data *BaseSysRoleModel.InfoResponse, err error) {

	err = BaseSysRoleModel.M.Where("id=?", id).Scan(&data)
	roleMenuList, err := BaseSysRoleMenuModel.M.Where("roleId", data.Id).Fields("menuId").All()
	for _, v := range roleMenuList {
		data.MenuIdList = append(data.MenuIdList, gconv.Int(v["menuId"]))
	}

	roleDepartmentList, err := BaseSysRoleDepartmentModel.M.Where("roleId", data.Id).Fields("departmentId").All()

	for _, v := range roleDepartmentList {
		data.DepartmentIdList = append(data.DepartmentIdList, gconv.Int(v["departmentId"]))
	}

	return
}

//添加操作
func (s *baseSysRoleService) Add(req *BaseSysRoleModel.AddReqParams) (insId int64, err error) {

	r, err := BaseSysRoleModel.M.OmitEmpty().Data(req).Insert()
	if err != nil {
		return 0, err
	}
	// 判断新增还是修改
	id, e := r.LastInsertId()
	//role 菜单关系绑定
	roleMenu := garray.NewArray()
	for _, v := range req.MenuIdList {
		roleMenu.Append(g.Map{
			"roleId": id,
			"menuId": v,
		})
	}
	BaseSysRoleMenuModel.M.OmitEmpty().Data(roleMenu).Insert()

	//role 部门绑定
	roleDepartment := garray.NewArray()
	for _, v := range req.DepartmentIdList {
		roleDepartment.Append(g.Map{
			"roleId":       id,
			"departmentId": v,
		})
	}
	BaseSysRoleDepartmentModel.M.OmitEmpty().Data(roleDepartment).Insert()
	return id, e
}

//修改操作
func (s *baseSysRoleService) Update(req *BaseSysRoleModel.UpdateReqParams) (Id int, err error) {

	_, err = BaseSysRoleModel.M.OmitEmpty().Data(req).Where("id=?", req.Id).Update()
	if err != nil {
		return 0, err
	}
	id := gconv.Int(req.Id)
	//role 菜单关系绑定
	roleMenu := garray.NewArray()
	for _, v := range req.MenuIdList {
		roleMenu.Append(g.Map{
			"roleId": req.Id,
			"menuId": v,
		})
	}
	//先删掉
	BaseSysRoleMenuModel.M.Where("roleId", id).Delete()
	BaseSysRoleMenuModel.M.OmitEmpty().Data(roleMenu).Insert()

	//role 部门绑定
	roleDepartment := garray.NewArray()
	for _, v := range req.DepartmentIdList {
		roleDepartment.Append(g.Map{
			"roleId":       id,
			"departmentId": v,
		})
	}
	//先删掉
	BaseSysRoleDepartmentModel.M.Where("roleId", id).Delete()
	BaseSysRoleDepartmentModel.M.OmitEmpty().Data(roleDepartment).Insert()
	return id, err
}

//分页查询
func (s *baseSysRoleService) Page(req *BaseSysRoleModel.PageReqParams) (total, page int, size int, list []*BaseSysRoleModel.Entity, err error) {
	page = req.Page
	size = req.Size

	M := BaseSysRoleModel.M.Where(1)

	if req.KeyWord != "" {
		M = M.WhereLike("name", "%"+req.KeyWord+"%")

	}
	if req.StartTime != "" {
		M = M.WhereGTE("createTIme", req.StartTime)
	}
	if req.EndTime != "" {
		M = M.WhereLTE("createTIme", req.EndTime)
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
	list = make([]*BaseSysRoleModel.Entity, len(data))
	err = data.Structs(&list)

	return
}

//List
func (s *baseSysRoleService) List(condition g.Map) (list []*BaseSysRoleModel.Entity, err error) {

	err = BaseSysRoleModel.M.Where(1).Where(condition).Scan(&list)
	return

}

//删除
func (s *baseSysRoleService) Delete(ids []int) (err error) {
	_, err = BaseSysRoleModel.M.Where("id IN(?)", ids).Delete()
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("删除模型数据失败")
		return
	}
	return
}

//权限数组 ["baseSysMenu:delete","baseSysMenu:update",...]
func (s *baseSysRoleService) GetPerms(roleIds string) (perms []string, err error) {
	var RoleMenuEntity []*BaseSysRoleMenuModel.Entity
	var MenuEntity []*BaseSysMenuModel.Entity
	roleIdsArr := gstr.Split(roleIds, ",")
	isSuperRole := false
	for _, v := range roleIdsArr {
		if v == "1" {
			isSuperRole = true
		}
	}
	if isSuperRole {
		err = BaseSysMenuModel.M.Scan(&MenuEntity)
		for _, m := range MenuEntity {
			permsArr := gstr.Split(m.Perms, ",")
			perms = append(perms, permsArr...)
		}
		perms = s.filerEmpty(perms)
		return
	}

	//否则根据具体角色查询相应的权限
	data, err := BaseSysRoleMenuModel.M.WhereIn("roleId", roleIdsArr).All()

	gconv.Structs(data, &RoleMenuEntity)

	var menuIds []int64
	for _, v := range RoleMenuEntity {
		menuIds = append(menuIds, v.MenuId)
	}

	menuData, err := BaseSysMenuModel.M.WhereIn("id", menuIds).All()
	gconv.Structs(menuData, &MenuEntity)

	for _, m := range MenuEntity {
		permsArr := gstr.Split(m.Perms, ",")
		perms = append(perms, permsArr...)
	}
	perms = s.filerEmpty(perms)
	return perms, nil

}

func (s *baseSysRoleService) GetMenus(roleIds string) (MenuEntity []*BaseSysMenuModel.Entity, err error) {
	var RoleMenuEntity []*BaseSysRoleMenuModel.Entity
	roleIdsArr := gstr.Split(roleIds, ",")
	isSuperRole := false
	for _, v := range roleIdsArr {
		if v == "1" {
			isSuperRole = true
		}
	}
	if isSuperRole {
		err = BaseSysMenuModel.M.OrderAsc("orderNum").Scan(&MenuEntity)

		return
	}
	//否则根据具体角色查询相应的权限
	err = BaseSysRoleMenuModel.M.WhereIn("roleId", roleIdsArr).Scan(&RoleMenuEntity)

	var menuIds []int64
	for _, v := range RoleMenuEntity {
		menuIds = append(menuIds, v.MenuId)
	}

	err = BaseSysMenuModel.M.WhereIn("id", menuIds).OrderAsc("orderNum").Scan(&MenuEntity)
	return
}

//过滤字符串空的
func (s *baseSysRoleService) filerEmpty(arr []string) []string {
	var res []string
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

//处理perms驼峰
func (s *baseSysRoleService) HandleMenus2Camel() (err error) {
	var MenuEntity []*BaseSysMenuModel.Entity

	err = BaseSysMenuModel.M.OrderAsc("orderNum").Scan(&MenuEntity)
	for _, m := range MenuEntity {
		if m.Perms != "" {
			var newPermsArr = garray.NewArray()

			permsArr := gstr.Split(m.Perms, ",")
			for _, p := range permsArr {
				if p != "" {
					newPermsArr.Append(s.doPerms2Camel(p))
				}
			}

			newPermsStr := newPermsArr.Join(",")
			//更新会以前的数据里
			BaseSysMenuModel.M.Where("id = ?", m.Id).Data(g.Map{"perms": newPermsStr}).Update()

		}

	}
	return
}

//应该是把base:sys:param:info, =》 baseSysParam:info,
//把base:app:im:message:read =》 baseAppImMessage:read
//把base:info, =》 base:info,
//把base, =》 base,
func (s *baseSysRoleService) doPerms2Camel(origin string) (p string) {
	arr := gstr.Split(origin, ":")
	sr := ":"
	if len(arr) <= 1 {
		return origin
	}
	for k, v := range arr {
		if len(arr)-1 != k {
			if k == 0 {
				p += gstr.LcFirst(v)
			} else {
				p += gstr.UcFirst(v)
			}
		} else {
			p += sr + v
		}
	}
	//if len(arr) == 4 {
	//	p = arr[0] + sr + gstr.LcFirst(arr[1]) + gstr.UcFirst(arr[2]) + sr + arr[3]
	//	return p
	//}
	//
	//if len(arr) == 5 {
	//	p = arr[0] + sr + gstr.LcFirst(arr[1]) + gstr.UcFirst(arr[2]) + sr + gstr.UcFirst(arr[3]) + sr + arr[4]
	//	return p
	//
	//}
	return
}
