package routers

import (
	"bdemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "Get:Get")

	// 用户
	beego.Router("/sys_user/login_out", &controllers.SysUserController{}, "Get:LogOut")
	beego.Router("/sys_user/login_form", &controllers.SysUserController{}, "Get:LoginForm")
	beego.Router("/sys_user/login_action", &controllers.SysUserController{}, "Post:LoginAction")

	// 首页
	beego.Router("/sys_home/index", &controllers.SysHomeController{}, "Get:Index")
	beego.Router("/sys_home/serverInfo", &controllers.SysHomeController{}, "Get:ServerInfo")

	// 菜单
	beego.Router("/sys_menu/list_sysmenu", &controllers.SysMenuController{}, "Get:ListSysMenu")
	beego.Router("/sys_menu/form_add_sysmenu", &controllers.SysMenuController{}, "Get:FormAddSysMenu")
	beego.Router("/sys_menu/form_modify_sysmenu", &controllers.SysMenuController{}, "Get:FormModifySysMenu")
	beego.Router("/sys_menu/save_sysmenu", &controllers.SysMenuController{}, "Post:SaveSysMenu")
	beego.Router("/sys_menu/add_sysmenu", &controllers.SysMenuController{}, "Post:AddSysMenu")
	beego.Router("/sys_menu/modify_sysmenu_status", &controllers.SysMenuController{}, "Get:ModifySysMenuStatus")
	beego.Router("/sys_menu/delete_sysmenu", &controllers.SysMenuController{}, "Get:DeleteSysMenu")

	// 角色
	beego.Router("/sys_role/list_sysrole", &controllers.SysRoleController{}, "Get:GetSysRoleListByPage")
	beego.Router("/sys_role/form_add_sysrole", &controllers.SysRoleController{}, "Get:FormAddSysRole")
	beego.Router("/sys_role/form_modify_sysrole", &controllers.SysRoleController{}, "Get:FormModifySysRole")
	beego.Router("/sys_role/save_sysrole", &controllers.SysRoleController{}, "Post:SaveSysRole")
	beego.Router("/sys_role/add_sysrole", &controllers.SysRoleController{}, "Post:AddSysRole")
	beego.Router("/sys_role/modify_sysrole_status", &controllers.SysRoleController{}, "Get:ModifySysRoleStatus")
	beego.Router("/sys_role/delete_sysrole", &controllers.SysRoleController{}, "Get:DeleteSysRole")

	// 日志
	beego.Router("/sys_log/list_syslog", &controllers.SysLogController{}, "Get:GetSysLogListByPage")
	beego.Router("/sys_log/delete_syslog", &controllers.SysLogController{}, "Get:DeleteSysLog")
}
